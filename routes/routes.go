package routes

import (
	"example/hello/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/api/register", controller.Register)

	return router
}
