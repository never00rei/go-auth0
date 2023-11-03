package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/never00rei/go-auth0/internal/models"
)

type RestHttpResponse struct {
	response http.Response
}

type RestResponseHander interface {
}

func NewRestResponse(response http.Response) *RestHttpResponse {
	return &RestHttpResponse{
		response: response,
	}
}

func (r RestHttpResponse) ProcessResponse(model models.Model) (models.Model, error) {
	if r.response.Body != nil {
		defer r.response.Body.Close()
	}

	// Add better error handling here...
	if r.response.StatusCode != 200 {
		return nil, fmt.Errorf("Request failed, response code: %d", r.response.StatusCode)
	}

	bodyResponseBytes, err := io.ReadAll(r.response.Body)
	if err != nil {
		return nil, err
	}

	modelErr := model.UnmarshalToModel(bodyResponseBytes)
	if modelErr != nil {
		return nil, modelErr
	}

	return model, nil
}
