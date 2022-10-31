package middleware

import "github.com/gin-gonic/gin"

func Monitor() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()
	}
}
