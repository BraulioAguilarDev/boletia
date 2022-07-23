package postgress

import (
	"boletia/model"
	"database/sql"
)

type CurrencyRepository struct {
	db *model.Queries
}

func NewCurrencyRepository(db *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{
		db: model.New(db),
	}
}

func (c CurrencyRepository) Get() error {
	// q := c.db.CreateCurrency()
	return nil
}

func (c CurrencyRepository) Pull() error {

	return nil
}
