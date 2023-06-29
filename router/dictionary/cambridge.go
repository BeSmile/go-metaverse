package dictionary

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/apis/dictionary"
)

func RegisterCambridgeRouter(v1 *gin.RouterGroup) {

	cambridgeGroup := v1.Group("/cambridge")

	cambridgeGroup.GET("/:word", dictionary.GetExplain)
}
