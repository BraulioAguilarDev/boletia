package currency

import "boletia/monitor"

type Usecase interface {
	Get() error
	Create(data monitor.Response) error
}
