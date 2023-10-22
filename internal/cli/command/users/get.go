package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/never00rei/go-auth0/internal/config"
	"github.com/never00rei/go-auth0/internal/models"

	"github.com/urfave/cli/v2"
)

const (
	endpoint = "users"
)

func GetAllUsers(c *cli.Context) error {

	tenant := os.Getenv(config.EnvSessionTenant)
	apiUrl := os.Getenv(config.EnvSessionApiUrl)
	sessionToken := os.Getenv(config.EnvSessionBearerToken)
	apiEndpoint := fmt.Sprintf("%s%s", apiUrl, endpoint)

	if os.Getenv("DEBUG") == "true" {
		log.Printf("Getting all users for %s tenant via: %s", tenant, apiEndpoint)
	}

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return fmt.Errorf("Failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", sessionToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to send request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Request failure, response not Http 200: %v", err)
	}

	var users []models.UserDetails

	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		return fmt.Errorf("Failed to decode JSON response: %v", err)
	}

	encodedUserData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to marshal user data: %v", err)
	}

	// Return the user data back to STDOUT So that it can be used.
	fmt.Println(string(encodedUserData))

	return nil
}
