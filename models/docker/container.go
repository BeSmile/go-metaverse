package docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"go-metaverse/helpers/app"
)

func (cmd *CMD) ContainerList(c *gin.Context) {
	containers, err := cmd.client.ContainerList(cmd.ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}
	fmt.Println("%s\n", containers)
	app.OK(c, containers, "success")
}

func (cmd *CMD) StopContainer(c *gin.Context) {
	id := c.Param("id")
	err := cmd.client.ContainerStop(cmd.ctx, id, nil)
	if err != nil {
		panic(err)
	}
	app.OK(c, nil, "success")
}

func (cmd *CMD) RmContainer(c *gin.Context) {
	id := c.Param("id")
	err := cmd.client.ContainerRemove(cmd.ctx, id, types.ContainerRemoveOptions{RemoveVolumes: true, Force: true})
	if err != nil {
		panic(err)
	}
	app.OK(c, nil, "success")
}

func (cmd *CMD) StartContainer(c *gin.Context) {
	id := c.Param("id")
	err := cmd.client.ContainerStart(cmd.ctx, id, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	app.OK(c, nil, "success")
}
