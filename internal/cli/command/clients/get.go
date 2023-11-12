package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/never00rei/go-auth0/internal/config"
	"github.com/never00rei/go-auth0/internal/handler"
	"github.com/never00rei/go-auth0/internal/models"
	"github.com/urfave/cli/v2"
)

const (
	endpoint = "clients"
)

func GetAllClients(c *cli.Context) error {

	tenant := os.Getenv(config.EnvSessionTenant)
	apiUrl := os.Getenv(config.EnvSessionApiUrl)
	sessionToken := os.Getenv(config.EnvSessionBearerToken)
	apiEndpoint := fmt.Sprintf("%s%s", apiUrl, endpoint)

	if os.Getenv("DEBUG") == "true" {
		log.Printf("Getting all clients for %s tenant via: %s", tenant, apiEndpoint)
	}

	var headers = [][]string{
		{"Accept", "application/json"},
		{"Authorization", sessionToken},
	}

	httpClient := http.DefaultClient

	httpRequest := handler.NewRestHttpRequest(apiEndpoint, httpClient)

	res, err := httpRequest.GetRequestHandler(headers)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Request failure, response not Http 200: %v", err)
	}

	var clients []models.Client

	if err := json.NewDecoder(res.Body).Decode(&clients); err != nil {
		return fmt.Errorf("failed to decode JSON response: %v", err)
	}

	encodedClientData, err := json.MarshalIndent(clients, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to marshal client data: %v", err)
	}

	fmt.Println(string(encodedClientData))

	return nil
}
