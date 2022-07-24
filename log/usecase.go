package log

import (
	"time"
)

type Usecase interface {
	Create(total time.Duration, code int, url string, tm time.Time) error
}
