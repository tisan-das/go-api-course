package middleware

import "github.com/gin-gonic/gin"

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
