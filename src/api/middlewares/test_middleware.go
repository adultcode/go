package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		apikey := c.GetHeader("X-Test")

		if apikey == "1" {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})

	}

}
