package handler

import (
	"boletia/currency"

	"github.com/gofiber/fiber/v2"
)

// MakeCurrencyHandler handler constructor
func MakeCurrencyHandler(usecase currency.Usecase, app *fiber.App) {
	currencyHandler := NewHandler(usecase)
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/currencies/:currency", currencyHandler.GetCurrencies)
}
