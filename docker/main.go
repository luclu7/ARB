package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		buildPackage(os.Args[1])
		fmt.Println("Your package is building, please check later...")
	} else {
		fmt.Println("Please provide an argument.")
		os.Exit(1)

	}
}
