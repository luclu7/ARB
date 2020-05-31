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
	"github.com/google/uuid"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func buildPackage(packageName string, uuid uuid.UUID) {
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
	s3host := getEnv("S3_HOST", "localhost")
	s3bucket := getEnv("S3_BUCKET", "arb")
	s3region := getEnv("S3_REGION", "us-east-1")

	var command strslice.StrSlice
	command = strslice.StrSlice(append([]string{"/bin/bash", "-c"}, "/build-aur "+packageName))

	envVars := []string{"MAIN_HOST=" + server, "S3_HOST=" + s3host, "S3_BUCKET=" + s3bucket, "S3_KEY=" + os.Getenv("S3_KEY"), "S3_SECRET=" + os.Getenv("S3_SECRET"), "S3_REGION=" + s3region}
	fmt.Println(envVars)
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
