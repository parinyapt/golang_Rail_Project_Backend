package middleware

import "github.com/gin-gonic/gin"

func Demoauth1(c *gin.Context) {
	if c.Query("pass1") == "1" {
		c.AbortWithStatus(403)
	}
	c.Next()
}

func Demoauth2(c *gin.Context) {
	if c.Query("pass2") == "2" {
		c.AbortWithStatus(403)
	}
	c.Next()
}