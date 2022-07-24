package postgress

import (
	"boletia/model"
	"boletia/monitor"
	"context"
	"database/sql"
	"fmt"
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

func (c CurrencyRepository) Get() error {
	// q := c.db.CreateCurrency()
	return nil
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
