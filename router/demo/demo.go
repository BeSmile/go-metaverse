package demo

import (
	"github.com/gin-gonic/gin"
)
func DemoBaseRouter(g *gin.RouterGroup) {
	v1 := g.Group("/api/v1")

	demoGroup := v1.Group("/demo")

	RegisterTemplateRouter(demoGroup)
}
