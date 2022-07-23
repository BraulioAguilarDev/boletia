package main

import (
	"boletia/server"
)

// main function
func main() {
	app, err := server.NewApp()
	if err != nil {
		panic(err)
	}

	// Execute sync
	go app.Sync()

	// fmt.Println(config.Config.ApiURL)

	// Run service
	if err := app.Run(":9090"); err != nil {
		panic(err)
	}
}
