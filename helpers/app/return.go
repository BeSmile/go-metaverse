package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOk())
}

func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnError(code))
}

func PageOk(c *gin.Context, data interface{}, count int, pageIndex int, pageSize int, msg int) {
	var res Response
	res.Data = data
}

//Normal 普通函数
func Normal(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}
