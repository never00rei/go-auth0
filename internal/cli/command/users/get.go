package users

import (
	"fmt"
	"log"

	"github.com/never00rei/go-auth0/internal/cli/shell"
	"github.com/never00rei/go-auth0/internal/config"
)

const (
	HEADERS  = "Accept: application/json"
	ENDPOINT = "/users"
)

func GetAllUsers() {
	if shell.CheckEnvVarExists(config.EnvSessionTenant) {
		fmt.Errorf("Tenant environment variable: %s does not exist, have you logged in?", config.EnvSessionTenant)
	}

	if shell.CheckEnvVarExists(config.EnvSessionBearerToken) {
		fmt.Errorf("Token environment variable: %s does not exist, have you logged in?", config.EnvSessionBearerToken)
	}

}
