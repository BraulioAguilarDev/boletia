package server

import (
	"boletia/pkg/monitor"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// App struct
type App struct {
	Service *fiber.App
	Period  int
}

// NewApp build a new instance
func NewApp() (*App, error) {
	// start db conexion
	// db...

	fiber := fiber.New()

	fiber.Use(cors.New())
	fiber.Use(logger.New())

	return &App{
		Service: fiber,
		Period:  5,
	}, nil
}

// SetupRouter makes routes
func (app *App) SetupRouter() {
	fmt.Println("router")
}

// Run function that execute api
func (app *App) Run(port string) error {

	app.SetupRouter()

	// Calls listen
	if err := app.Service.Listen(port); err != nil {
		return err
	}

	return nil
}

// Sync function calls internal function to get infomation
func (app *App) Sync() {
	currencyMonitor := monitor.NewHandler(app.Period)
	for {
		<-time.After(time.Duration(app.Period) * time.Second)
		res, err := currencyMonitor.Pull()
		if err != nil {
			panic(err)
		}

		fmt.Println(res.Status)
		// Passing the result to new function for keeping it in our db
	}
}
