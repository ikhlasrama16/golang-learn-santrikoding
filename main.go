package main

import (
	"example/hello/config"
	"example/hello/database"
	"example/hello/routes"
)

func main() {

	config.LoadEnv()

	//inisialisasi database
	database.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))

}
