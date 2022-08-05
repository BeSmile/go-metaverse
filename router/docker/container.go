package docker

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/apis/docker"
)

func RegisterDockerContainerRoute(v1 *gin.RouterGroup) {
	containerRouter := v1.Group("/container")

	v1.GET("/containers", docker.GetAllContainers)

	containerRouter.PUT("/:id", docker.StopContainer)
	containerRouter.DELETE("/:id", docker.RmContainer)
	containerRouter.POST("/:id", docker.StartContainer)
}
