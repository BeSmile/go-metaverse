package middleware

import (
	"github.com/gin-gonic/gin"
)

// InitMiddleware 初始化中间件
func InitMiddleware(r *gin.Engine) {
	// 跨域处理
	r.Use(Options)
	// 缓存处理
	r.Use(NoCache)
	r.Use(Secure)
	// 生成requestId
	r.Use(RequestId())
	r.Use(LoggerToFile()) // 返回一个函数
}
