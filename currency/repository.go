package currency

import "boletia/monitor"

type Repository interface {
	Get() error
	Create(data monitor.Response) error
}
