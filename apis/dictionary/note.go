package dictionary

import (
	"github.com/gin-gonic/gin"
	"go-metaverse/helpers/app"
	"go-metaverse/models/dictionary"
)

func InsertNote(c *gin.Context) {
	var note dictionary.Note
	if err := c.Bind(&note); err !=nil {
		app.Error(c, 400, err, "")
		return
	}
	if ierror := note.Insert(); ierror != nil{
		app.Error(c, 500, ierror, "")
		return
	}
	app.OK(c, note, "")
}
