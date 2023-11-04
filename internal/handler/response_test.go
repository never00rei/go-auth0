package handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/never00rei/go-auth0/internal/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockModel is a mock implementation of the Model interface for testing.
type MockModel struct {
	mock.Mock
}

func (m *MockModel) UnmarshalToModel(data []byte) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestProcessResponse_Success(t *testing.T) {
	// Arrange
	jsonResponse := `{"key":"value"}`
	mockResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(jsonResponse)),
	}

	mockModel := new(MockModel)
	mockModel.On("UnmarshalToModel", mock.Anything).Return(nil)

	restResponse := handler.NewRestResponse(mockResponse)

	// Act
	resultModel, err := restResponse.ProcessResponse(mockModel)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultModel)
	mockModel.AssertExpectations(t) // Assert that UnmarshalToModel was called
}
