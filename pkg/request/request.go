package request

import (
	"bytes"
	"net/http"
	"time"
)

// Global response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

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

// Success returns a simple payload for success
func Success(msg string) Response {
	return Response{
		Success: true,
		Message: msg,
	}
}

// Failure returns a simple payload for failed
func Failure(msg string) Response {
	return Response{
		Success: false,
		Message: msg,
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
