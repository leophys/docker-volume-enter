package main

import (
	"context"
	"fmt"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

func getLocalPath(mount types.MountPoint) string {
	switch mount.Type {
	case "bind":
		fmt.Println("bind mount")
	case "volume":
		fmt.Println("volume mount")
	}
	return mount.Source
}

func main() {
	client, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		for _, mount := range container.Mounts {
			fmt.Printf("%s %s\n", container.ID[:10], getLocalPath(mount))
		}
	}
}
