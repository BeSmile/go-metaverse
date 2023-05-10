package router

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	middleware.InitMiddleware(r) // 添加中间件

	InitDomainRouter(r)

	return r
}
