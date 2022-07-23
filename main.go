package main

import (
	"boletia/pkg/monitor"
	"boletia/server"
	"fmt"
)

// main function
func main() {
	app, err := server.NewApp()
	if err != nil {
		panic(err)
	}

	// ====================
	//  Moving to goroutine
	currencyMonitor := monitor.NewHandler(app.Period)
	result, err := currencyMonitor.Sync()
	if err != nil {
		panic(err)
	}
	fmt.Println(result.StatusCode)
	// =====================

	// Run service
	if err := app.Run(":9090"); err != nil {
		panic(err)
	}
}
