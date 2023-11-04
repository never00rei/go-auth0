package handler

import (
	"fmt"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RestHttpRequest struct {
	url    string
	client HttpClient
}

type RestRequestHandler interface {
	GetRequestHandler()
	NewRestHttpRequest()
}

func NewRestHttpRequest(url string, client HttpClient) *RestHttpRequest {

	if client == nil {
		client = http.DefaultClient
	}

	return &RestHttpRequest{
		url:    url,
		client: client,
	}
}

func (r RestHttpRequest) GetRequestHandler(headers [][]string) (*http.Response, error) {

	req, err := http.NewRequest("GET", r.url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for _, header := range headers {
		req.Header.Add(header[0], header[1])
	}

	res, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	return res, nil
}
