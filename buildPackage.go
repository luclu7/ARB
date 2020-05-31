package main

import (
	"context"
	log "github.com/sirupsen/logrus"
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

func buildPackage(packageName string, uuid string, secret string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	imageName := "docker.io/luclu7/docker-arch-build-aur:latest"
	// don't care about the output
	_, err = cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	server := getEnv("MAIN_HOST", "localhost")
	s3host := getEnv("S3_HOST", "localhost")
	s3bucket := getEnv("S3_BUCKET", "arb")
	s3region := getEnv("S3_REGION", "us-east-1")

	var command strslice.StrSlice
	command = strslice.StrSlice(append([]string{"/bin/bash", "-c"}, "/build-aur "+packageName))

	envVars := []string{"MAIN_HOST=" + server, "S3_HOST=" + s3host, "S3_BUCKET=" + s3bucket, "S3_KEY=" + os.Getenv("S3_KEY"), "S3_SECRET=" + os.Getenv("S3_SECRET"), "S3_REGION=" + s3region, "UUID=" + uuid, "SECRET=" + secret}
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

	log.WithFields(log.Fields{
		"package":   packageName,
		"UUID":      uuid,
		"container": resp.ID,
	}).Info("New container")

}
