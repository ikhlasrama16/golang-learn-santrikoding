package main

import (
	"example/hello/config"
	"example/hello/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	//inisialisasi database
	database.InitDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello WOrld",
		})
	})

	router.Run(":" + config.GetEnv("APP_PORT", "3000"))

}
