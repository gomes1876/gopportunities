package router

import "github.com/gin-gonic/gin"

func Initialize() {

	routes := gin.Default()

	initializeRoutes(routes)

	routes.Run(":8080")
}
