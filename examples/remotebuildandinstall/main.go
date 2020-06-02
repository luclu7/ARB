package main

import (
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	resty "github.com/go-resty/resty/v2"
	ARB "github.com/luclu7/ARB/API"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func waitForBuildComplete(uuid string, server string, client *resty.Client) bool {
	status, err := ARB.GetStatusOfBuild(uuid, server, client)
	if err != nil {
		panic(err)
	}
	if status.Status != 1 {
		return false
	} else {
		return true
	}
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Please provide a package to install.")
		os.Exit(1)
	}
	spinner := spinner.New(spinner.CharSets[14], 50*time.Millisecond)
	spinner.Start()

	var server string
	var pkg string

	flag.StringVar(&server, "server", "http://localhost:8081", "server address:port")
	flag.StringVar(&pkg, "pkg", "xcwd-git", "package to build")
	flag.Parse()

	client := resty.New()
	RequestResponse, err := ARB.LaunchBuild(pkg, server, client)
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second) // wait for initial launch
	for waitForBuildComplete(RequestResponse.UUID, server, client) != true {
		time.Sleep(5 * time.Second)
	}
	URLs, err := ARB.GetURLs(RequestResponse.UUID, server, client)
	fmt.Println("Build finished! Download...")
	var packagesToInstall string
	for _, file := range URLs {
		fileName := filepath.Base(file.URL)
		err := downloadFile(fileName, file.URL)
		if err != nil {
			panic(err)
		}
		packagesToInstall = packagesToInstall + " " + fileName
	}
	cmd := exec.Command("/bin/sh", "-c", "sudo pacman --noconfirm -U "+packagesToInstall)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	spinner.Stop()
}
