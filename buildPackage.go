package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func buildPackage(packageName string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	imageName := "docker.io/luclu7/docker-arch-build-aur:latest"

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	server := getEnv("MAIN_HOST", "localhost")

	var command strslice.StrSlice
	command = strslice.StrSlice(append([]string{"/bin/bash", "-c"}, "/build-aur "+packageName+"; curl "+server+"/build/complete/mark/"+packageName))

	envVars := []string{"MAIN_HOST=10.10.5.42:8081"}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   command,
		Env:   envVars,
	},
		nil,
		nil,
		"")

	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}
