package docker

import (
	"github.com/gin-gonic/gin"
	docker "go-metaverse/models/docker"
)

func GetAllImages(c *gin.Context) {
	docker.Cmd.ImageList(c)
}
