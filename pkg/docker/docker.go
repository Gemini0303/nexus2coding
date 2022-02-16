package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err.Error())
	}
	reader, err := cli.ImagePull(context.Background(),
		"43.254.44.133:58000/repository/coding-private-proxy/v2/coding-private/release/coding-cd/manifests/5.0.0-20220214-091242-a73595", types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	reader.Close()
}
