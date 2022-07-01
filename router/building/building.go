package building

import "github.com/gin-gonic/gin"

func InitBuildingRouter(g *gin.RouterGroup) {
	v1 := g.Group("/api/v1")

	buildingGroup := v1.Group("/building")

	InitMaterialRouter(buildingGroup);
}