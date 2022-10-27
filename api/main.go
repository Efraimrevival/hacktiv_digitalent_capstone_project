package main

import (
	"api/src/config/env"
	"api/src/database"
	"api/src/router"
)

func main() {
	database.InitPostgres()
	router.StartServer().Run(":" + env.GetServerEnv())
}
