package usecase

import (
	"boletia/currency"
	"boletia/monitor"
)

type CurrencyUsecase struct {
	CurrencyRepo currency.Repository
}

func NewCurrencyUsecase(repo currency.Repository) *CurrencyUsecase {
	return &CurrencyUsecase{
		CurrencyRepo: repo,
	}
}

func (c CurrencyUsecase) Create(r monitor.Response) error {
	// Calls repository function to save
	if err := c.CurrencyRepo.Create(r); err != nil {
		return err
	}

	return nil
}
