package postgress

import (
	"boletia/model"
	"boletia/monitor"
	"context"
	"database/sql"
	"time"

	"github.com/golang/glog"
)

type CurrencyRepository struct {
	db  *model.Queries
	ctx context.Context
}

func NewCurrencyRepository(db *sql.DB, ctx context.Context) *CurrencyRepository {
	return &CurrencyRepository{
		db:  model.New(db),
		ctx: ctx,
	}
}

// Create function to insert rows from api
func (c CurrencyRepository) Create(data monitor.Response) error {

	// TODO: check the best option for this
	for _, item := range data.Data {
		currency, err := c.db.CreateCurrency(c.ctx, model.CreateCurrencyParams{
			Code:      item.Code,
			Value:     item.Value,
			UpdatedAt: data.Meta.LastUpdatedAt,
		})

		glog.Infof("Saved: %v", currency)

		if err != nil {
			return err
		}
	}

	glog.Infof("%d rows was inserted.", len(data.Data))

	return nil
}

// GetCurrenciesByCode retuns a currencies list filtered by code
func (c CurrencyRepository) GetCurrenciesByCode(code string, start, end time.Time) ([]model.Currency, error) {
	data, err := c.db.GetCurrenciesByCode(c.ctx, model.GetCurrenciesByCodeParams{
		Code:        code,
		UpdatedAt:   start,
		UpdatedAt_2: end,
	})

	return data, err
}

// GetAllCurrencies retuns all rows from db
func (c CurrencyRepository) GetAllCurrencies(start, end time.Time) ([]model.Currency, error) {
	data, err := c.db.GetAllCurrencies(c.ctx, model.GetAllCurrenciesParams{
		UpdatedAt:   start,
		UpdatedAt_2: end,
	})

	return data, err
}
