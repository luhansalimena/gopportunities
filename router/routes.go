package router

import "github.com/gin-gonic/gin"

func InitliazeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
