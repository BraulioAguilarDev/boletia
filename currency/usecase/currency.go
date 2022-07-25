package usecase

import (
	"boletia/currency"
	"boletia/model"
	"boletia/monitor"
	"time"
)

type CurrencyUsecase struct {
	CurrencyRepo currency.Repository
}

// NewCurrencyUsecase generates a new instance
func NewCurrencyUsecase(repo currency.Repository) *CurrencyUsecase {
	return &CurrencyUsecase{
		CurrencyRepo: repo,
	}
}

// Create is a usecases that calls to create function in the repository pkg
func (c CurrencyUsecase) Create(r monitor.Response) error {
	// Calls repository function to save
	if err := c.CurrencyRepo.Create(r); err != nil {
		return err
	}

	return nil
}

// GetCurrenciesByCode is a function to get a currencies list -- this calls to repository
func (c CurrencyUsecase) GetCurrenciesByCode(code string, start, end time.Time) ([]currency.Currency, error) {

	result, err := c.CurrencyRepo.GetCurrenciesByCode(code, start, end)

	// maps internal to external data
	data := mapModelCurrencyToDto(result)

	return data, err
}

// GetAllCurrencies function gets ALL rows
func (c CurrencyUsecase) GetAllCurrencies(start, end time.Time) ([]currency.Currency, error) {

	result, err := c.CurrencyRepo.GetAllCurrencies(start, end)

	// maps internal to external data
	data := mapModelCurrencyToDto(result)

	return data, err
}

// mapModelCurrencyToDto function maps model struct to an external object
func mapModelCurrencyToDto(currencies []model.Currency) []currency.Currency {
	var list []currency.Currency
	for _, c := range currencies {
		list = append(list, currency.Currency{
			Code:  c.Code,
			Value: c.Value,
			Date:  c.UpdatedAt,
		})
	}

	return list
}
