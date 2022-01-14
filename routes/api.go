package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "world at v1",
			})
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "world at v2",
			})
		})
	}
}
