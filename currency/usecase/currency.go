package usecase

import (
	"boletia/currency"
)

type CurrencyUsecase struct {
	CurrencyRepo currency.Repository
}

func NewCurrencyUsecase(repo currency.Repository) *CurrencyUsecase {
	return &CurrencyUsecase{
		CurrencyRepo: repo,
	}
}

func (c CurrencyUsecase) Get() error {
	// c.CurrencyRepo.Get()
	return nil
}

func (c CurrencyUsecase) Pull() error {
	// c.CurrencyRepo.Bulk()
	return nil
}
