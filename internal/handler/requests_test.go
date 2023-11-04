package handler_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/never00rei/go-auth0/internal/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHttpClient is a mock for the HttpClient interface
type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetRequestHandler_Success(t *testing.T) {
	// Arrange
	mockClient := new(MockHttpClient)
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
	}

	mockClient.On("Do", mock.Anything).Return(mockResponse, nil)

	restRequest := handler.NewRestHttpRequest("http://example.com", mockClient)

	// Act
	response, err := restRequest.GetRequestHandler(nil)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	if response != nil {
		assert.Equal(t, http.StatusOK, response.StatusCode)
	}

	mockClient.AssertExpectations(t) // Verify that the mock was called
}

func TestGetRequestHandler_FailToSendRequest(t *testing.T) {
	// Arrange
	mockClient := new(MockHttpClient)

	// Because we're not expecting a repsonse, we're setting up a nil response typed
	// to *http.Response. This is to simulate a network failure.
	var nilResponse *http.Response
	mockClient.On("Do", mock.Anything).Return(nilResponse, fmt.Errorf("network error"))

	restRequest := handler.NewRestHttpRequest("http://example.com", mockClient)

	// Act
	response, err := restRequest.GetRequestHandler(nil)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.EqualError(t, err, "failed to send request: network error")

	mockClient.AssertExpectations(t)
}
