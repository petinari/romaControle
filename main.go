package main

import (
	"romaControle/config"
	"romaControle/router"
)

var (
	logger *config.Logger
)

func main() {

	//run database

	// Initialize Router
	router.Initialize()
}
