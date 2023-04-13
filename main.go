package main

import (
	"romaControle/config"
	"romaControle/router"
)

func init() {
	config.LoadEnv()
	config.InitMongo()
	config.InitPostgres()
	config.InitRedis()

}

func main() {

	//run database

	// Initialize Router
	router.Initialize()
}
