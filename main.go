package main

import (
	"boletia/server"
	"flag"

	"github.com/golang/glog"
)

// main function
func main() {
	app, err := server.NewApp()
	if err != nil {
		glog.Errorf("Initialization failed: %s", err.Error())
	}

	flag.Parse()

	// Execute sync with period time
	go app.Sync()

	// Run service
	if err := app.Run(":9090"); err != nil {
		glog.Errorf("Run failed: %s", err.Error())
	}
}
