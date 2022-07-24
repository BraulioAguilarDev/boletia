package monitor

import (
	"boletia/config"
	"boletia/pkg/request"
	"net/http"
	"time"
)

// Code object by each currency
type Code struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

// Meta it´s a metadata
type Meta struct {
	LastUpdatedAt time.Time `json:"last_updated_at"`
}

// Response represents an object that we receives from external api
type Response struct {
	Meta Meta            `json:"meta"`
	Data map[string]Code `json:"data"`
}

type handler struct {
	Client *request.Request
}

func NewHandler() *handler {
	return &handler{
		Client: request.NewClient(config.Config.Timeout),
	}
}

// GetCurrencies requests external api
func (h *handler) GetCurrencies() (*http.Response, time.Time, error) {
	res, err := h.Client.MakeRequest(&request.Options{
		Method:   http.MethodGet,
		Endpoint: config.Config.ApiURL + "?apikey=" + config.Config.ApiKey,
	})

	if err != nil {
		return nil, time.Now(), err
	}

	return res, time.Now(), nil
}
