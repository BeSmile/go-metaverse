package docker

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/models/docker"
)

func GetAllContainers(c *gin.Context) {
	docker.Cmd.ContainerList(c)
}

func StopContainer(c *gin.Context) {
	docker.Cmd.StopContainer(c)
}

func RmContainer(c *gin.Context) {
	docker.Cmd.RmContainer(c)
}

func StartContainer(c *gin.Context) {
	docker.Cmd.StartContainer(c)
}
