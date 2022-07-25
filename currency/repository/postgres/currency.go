package postgress

import (
	"boletia/model"
	"boletia/monitor"
	"context"
	"database/sql"
	"fmt"
	"time"
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
		created, err := c.db.CreateCurrency(c.ctx, model.CreateCurrencyParams{
			Code:      item.Code,
			Value:     item.Value,
			UpdatedAt: data.Meta.LastUpdatedAt,
		})

		fmt.Println("inserted:", created)

		if err != nil {
			return err
		}
	}

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
