package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-metaverse/docs"
	"go-metaverse/handler"
	buildingRouter "go-metaverse/router/building"
	demoRouter "go-metaverse/router/demo"
)

func baseRouter(g *gin.RouterGroup) {
	g.GET("/info", handler.Ping)
}

func InitDomainRouter(r *gin.Engine) *gin.RouterGroup {
	dGroup := r.Group("")

	staticFileRouter(dGroup, r)

	baseRouter(dGroup)

	initDemoRouter(dGroup)

	initBuildingRouter(dGroup)

	// 注册swagger路由
	swaggerRouter(dGroup)

	return dGroup
}

// 初始化demoRouter
func initDemoRouter(g *gin.RouterGroup) {
	demoRouter.DemoBaseRouter(g) // 注入demo路由
}

func initBuildingRouter(g *gin.RouterGroup) {
	buildingRouter.InitBuildingRouter(g) // 注入demo路由
}

// 配置静态资源目录
func staticFileRouter(g *gin.RouterGroup, r *gin.Engine) {
	// 配置static静态目录
	g.Static("/static", "static")
}

// 配置swagger路由
func swaggerRouter(g *gin.RouterGroup) {
	fmt.Println(111)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
