package currency

import (
	"boletia/monitor"
)

type Usecase interface {
	Create(data monitor.Response) error
}
