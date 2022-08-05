package docker

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/apis/docker"
)

func RegisterDockerImageRouter(v1 *gin.RouterGroup) {
	v1.GET("/images", docker.GetAllImages)

	imageRouter := v1.Group("/image")

	imageRouter.GET("/list", docker.GetAllImages)
}
