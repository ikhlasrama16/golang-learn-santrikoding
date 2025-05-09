package routes

import (
	"example/hello/controller"
	"example/hello/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/api/register", controller.Register)

	router.POST("/api/login", controller.Login)

	router.GET("/api/users", middlewares.AuthMiddleware(), controller.FindUser)

	router.POST("/api/users", middlewares.AuthMiddleware(), controller.CreateUser)

	return router
}
