package currency

import (
	"boletia/model"
	"boletia/monitor"
	"time"
)

type Repository interface {
	Create(data monitor.Response) error
	GetCurrenciesByCode(code string, start, end time.Time) ([]model.Currency, error)
	GetAllCurrencies(start, end time.Time) ([]model.Currency, error)
}
