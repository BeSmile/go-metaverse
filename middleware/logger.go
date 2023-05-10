package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// LoggerToFile 记录文件  /**
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		statusCode := c.Writer.Status()

		reqUri := c.Request.RequestURI

		clientIP := c.ClientIP()

		// todo 实现输出日志文件
		fmt.Printf(" %s %3d %13v %15s %s %s",
			startTime.Format("2006-01-02 15:04:05.9999"),
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)

	}

}
