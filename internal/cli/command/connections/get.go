package connections

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
	endpoint = "connections"
)

func GetAllConnections(c *cli.Context) error {

	tenant := os.Getenv(config.EnvSessionTenant)
	apiUrl := os.Getenv(config.EnvSessionApiUrl)
	sessionToken := os.Getenv(config.EnvSessionBearerToken)
	apiEndpoint := fmt.Sprintf("%s%s", apiUrl, endpoint)

	if os.Getenv("DEBUG") == "true" {
		log.Printf("Getting all connections or %s tenant via: %s", tenant, apiEndpoint)
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
		return fmt.Errorf("Request Failure, response not Http 200: %v", err)
	}

	var connections []models.Connection

	if err := json.NewDecoder(res.Body).Decode(&connections); err != nil {
		return fmt.Errorf("Faile to decode JSON response: %v", err)
	}

	encodedConnectionData, err := json.MarshalIndent(connections, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to marshal connection data: %v", err)

	}

	fmt.Println(string(encodedConnectionData))

	return nil
}
