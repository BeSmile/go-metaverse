package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)

type CMD struct {
	client *client.Client
	ctx    context.Context
}

var Cmd CMD

func InitCmdBackendEnv() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	fmt.Println("InitCmdBackendEnv")
	Cmd = CMD{
		client: cli,
		ctx:    ctx,
	}
	fmt.Println(Cmd)
}
