package server

import (
	"boletia/config"
	"boletia/currency"
	postgress "boletia/currency/repository/postgres"
	usecaseCurrency "boletia/currency/usecase"
	"boletia/pkg/monitor"
	"database/sql"
	"fmt"
	"time"

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
		panic(err)
	}

	// Build repositories
	currencyRepo := postgress.NewCurrencyRepository(db)

	fiber := fiber.New()

	fiber.Use(cors.New())
	fiber.Use(logger.New())

	return &App{
		Service:         fiber,
		CurrencyUsecase: usecaseCurrency.NewCurrencyUsecase(currencyRepo),
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
	currencyMonitor := monitor.NewHandler(5) // period sync
	for {
		<-time.After(time.Duration(10) * time.Second)
		res, err := currencyMonitor.Pull()
		if err != nil {
			panic(err)
		}

		fmt.Println(res.Status)
		// Passing the result to new function for keeping it in our db
	}
}

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

	return db, err
}
