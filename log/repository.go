package log

import (
	"boletia/model"
)

type Repository interface {
	Create(data model.CreateLogParams) error
}
