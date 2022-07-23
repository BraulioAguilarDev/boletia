package monitor

import (
	"boletia/pkg/request"
	"net/http"
	"time"
)

type handler struct {
	Period time.Duration // remove period time to config yml
	Client *request.Request
}

func NewHandler(period int) *handler {
	return &handler{
		Period: time.Duration(period),
		Client: request.NewClient(5),
	}
}

// Pull requests external api
func (h *handler) Pull() (*http.Response, error) {
	res, err := h.Client.MakeRequest(&request.Options{
		Method:   http.MethodGet,
		Endpoint: "https://pokeapi.co/api/v2/pokemon/ditto",
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
