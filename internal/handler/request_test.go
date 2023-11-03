package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetRequestHandler tests the GetRequestHandler function
func TestGetRequestHandler(t *testing.T) {
	// Set up a test server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer test-token" {
			t.Errorf("Authorization header was not set correctly. Got %s, want %s", authHeader, "Bearer test-token")
		}
		io.WriteString(w, `{"status": "ok"}`)
	}))
	defer testServer.Close()

	restRequest := NewRestHttpRequest(testServer.URL)

	// Define the headers you want to send
	headers := [][]string{
		{"Authorization", "Bearer test-token"},
	}

	// Make the request using the GetRequestHandler
	resp, err := restRequest.GetRequestHandler(headers)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Expected no error reading response body, got %v", err)
	}

	expectedBody := `{"status": "ok"}`
	if string(body) != expectedBody {
		t.Errorf("Expected body to be %s, got %s", expectedBody, body)
	}
}
