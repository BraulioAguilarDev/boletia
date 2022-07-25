package currency

import (
	"boletia/monitor"
	"time"
)

type Usecase interface {
	Create(data monitor.Response) error
	GetCurrenciesByCode(code string, start, end time.Time) ([]Currency, error)
}
