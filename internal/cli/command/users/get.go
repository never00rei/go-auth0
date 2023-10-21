package users

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/never00rei/go-auth0/internal/config"

	//"github.com/never00rei/go-auth0/internal/models"
	"github.com/urfave/cli/v2"
)

const (
	endpoint = "/users"
)

func GetAllUsers(c *cli.Context) error {

	tenant := os.Getenv(config.EnvSessionTenant)
	apiUrl := os.Getenv(config.EnvSessionApiUrl)
	sessionToken := os.Getenv(config.EnvSessionBearerToken)
	apiEndpoint := fmt.Sprintf("%s/%s", apiUrl, endpoint)

	log.Printf("Fetching users from %s: %s", tenant, apiEndpoint)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return fmt.Errorf("Failed to create request: %v", err)
	}

	return nil
}
