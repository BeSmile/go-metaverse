package building

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/apis/building"
)

// 初始化素材岛路由
func InitMaterialRouter(g *gin.RouterGroup) {
	materialRouter := g.Group("/material")
	// 动物岛
	//materialRouter
	materialRouter.GET("/animals", building.GetMaterialLandAnimal)
	materialRouter.POST("/magic", building.SaveMagic)
}
