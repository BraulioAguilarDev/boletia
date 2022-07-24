package currency

import (
	"boletia/monitor"
)

type Repository interface {
	Create(data monitor.Response) error
}
