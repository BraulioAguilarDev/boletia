package main

import (
	"boletia/config"
	"boletia/server"
	"flag"
	"fmt"

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
	if err := app.Run(":" + fmt.Sprint(config.Config.Port)); err != nil {
		glog.Errorf("Run failed: %s", err.Error())
	}
}
