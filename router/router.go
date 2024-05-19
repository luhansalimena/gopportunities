package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()

	InitliazeRoutes(router)

	router.Run() // listen and serve on
}
