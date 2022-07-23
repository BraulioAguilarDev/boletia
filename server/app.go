package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// App struct
type App struct {
	Service *fiber.App
	Period  int
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
