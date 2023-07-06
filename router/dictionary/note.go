package dictionary

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/apis/dictionary"
)

func RegisterNoteRouter(v1 *gin.RouterGroup) {
	noteGroup := v1.Group("/note")
	noteGroup.POST("", dictionary.InsertNote)
}