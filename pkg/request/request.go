package request

import (
	"bytes"
	"net/http"
	"time"
)

// Request struct
type Request struct {
	Client *http.Client
}

// Options struct
type Options struct {
	Endpoint string
	Body     []byte
	Method   string
}

// NewClient func
func NewClient(timeout int) *Request {
	return &Request{
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

// MakeRequest func executes external request
// returns *http.Response and error
func (r *Request) MakeRequest(o *Options) (*http.Response, error) {
	req, err := http.NewRequest(o.Method, o.Endpoint, bytes.NewBuffer(o.Body))
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
