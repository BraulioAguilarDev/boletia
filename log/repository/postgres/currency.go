package postgress

import (
	"boletia/model"
	"context"
	"database/sql"
)

type LogRepository struct {
	db  *model.Queries
	ctx context.Context
}

func NewLogRepository(db *sql.DB, ctx context.Context) *LogRepository {
	return &LogRepository{
		db:  model.New(db),
		ctx: ctx,
	}
}

func (l LogRepository) Create(data model.CreateLogParams) error {

	_, err := l.db.CreateLog(l.ctx, data)

	if err != nil {
		return err
	}

	return nil
}
