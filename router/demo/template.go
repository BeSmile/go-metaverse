package demo

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/apis/template"
)

func RegisterTemplateRouter(v1 *gin.RouterGroup) {
	templateRouter := v1.Group("/templates")

	templateRouter.GET("/:id", template.GetTemplateById)
}
