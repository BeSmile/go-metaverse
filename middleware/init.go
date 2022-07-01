package middleware

import (
	"github.com/gin-gonic/gin"
)

// InitMiddleware 初始化中间件
func InitMiddleware(r *gin.Engine) {
	r.Use(Options)
	r.Use(NoCache)
	r.Use(Secure)
	r.Use(LoggerToFIle()) // 返回一个函数
}
