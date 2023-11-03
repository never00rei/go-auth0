package handler

import (
	"fmt"
	"net/http"
)

type RestHttpRequest struct {
	url string
}

type RestRequestHandler interface {
	GetRequestHandler()
	NewRestHttpRequest()
}

func NewRestHttpRequest(url string) *RestHttpRequest {
	return &RestHttpRequest{
		url: url,
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

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	return res, nil
}
