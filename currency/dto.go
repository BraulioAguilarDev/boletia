package currency

import "time"

type Currency struct {
	Code  string    `json:"code"`
	Value float64   `json:"value"`
	Date  time.Time `json:"updated"`
}
