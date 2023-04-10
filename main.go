package main

import (
	"romaControle/config"
	"romaControle/router"
)

func init() {
	config.LoadEnv()
	config.InitMongo()

}

func main() {

	//run database

	// Initialize Router
	router.Initialize()
}
