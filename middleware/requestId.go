package middleware

import (
	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
	"log"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-ID")

		if requestId == "" {
			u4 := uuid.Generate()
			requestId = u4.String()
		}
		log.Println("Request Id: ", requestId)
		c.Set("X-Request-ID", requestId)

		c.Writer.Header().Set("X-Request-ID", requestId)
		c.Next()
	}
}
