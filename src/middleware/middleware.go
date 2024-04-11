package middleware

import (
	"go-api-course/src/logging"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Authenticate(c *gin.Context) {
	if c.Request.Header.Get("Token") != "auth" {
		c.AbortWithStatusJSON(401, gin.H{
			"msg": "token not found",
		})
		return
	}
	c.Next()
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key2", "Value2")
	c.Next()
}

func SetRequestId(c *gin.Context) {
	requestId := uuid.New().String()
	c.Set("requestId", requestId)
	c.Next()
}

func LogRequest(logger logging.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		logger.LogRequest(c)
		c.Next()
	}
}

func LogResponse(logger logging.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		logger.LogResponse(c)
	}
}
