package server

import (
	"boletia/config"
	"boletia/currency"
	currencyRepository "boletia/currency/repository/postgres"
	usecaseCurrency "boletia/currency/usecase"
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

// App struct
type App struct {
	Service         *fiber.App
	CurrencyUsecase currency.Usecase
}

// NewApp build a new instance
func NewApp() (*App, error) {
	// start db conexion
	db, err := initDb()
	if err != nil {
		return nil, err
	}

	// Build instance repositories & app
	ctx := context.Background()
	currencyRepo := currencyRepository.NewCurrencyRepository(db, ctx)

	fiber := fiber.New()

	fiber.Use(cors.New())
	fiber.Use(logger.New())

	return &App{
		Service:         fiber,
		CurrencyUsecase: usecaseCurrency.NewCurrencyUsecase(currencyRepo),
	}, nil
}

// SetupRouter makes public routes
func (app *App) SetupRouter() {
	// fmt.Println("router")
}

// Run function executes our service
func (app *App) Run(port string) error {

	app.SetupRouter()

	if err := app.Service.Listen(port); err != nil {
		return err
	}

	return nil
}

// initDb function to do conexion to db
func initDb() (*sql.DB, error) {
	sourceName := fmt.Sprintf(`host=%s port=%d user=%s dbname=%s password=%s sslmode=%s`,
		config.Config.Database.DBhost,
		config.Config.Database.DBport,
		config.Config.Database.DBuser,
		config.Config.Database.DBname,
		config.Config.Database.DBpassword,
		config.Config.Database.SSLmode,
	)

	db, err := sql.Open(config.Config.Database.Driver, sourceName)

	// Check ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
