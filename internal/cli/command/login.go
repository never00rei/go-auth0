package command

import (
	"fmt"
	"log"

	"github.com/never00rei/go-auth0/internal/auth"
	"github.com/never00rei/go-auth0/internal/config"
	"github.com/urfave/cli/v2"
)

func Login(c *cli.Context) error {
	tenant := c.String("tenant")

	clientCreds, err := config.GetCredentials(tenant)
	if err != nil {
		return fmt.Errorf("failed to read credentials file: %v", err)
	}

	if clientCreds.ClientID == "" || clientCreds.ClientSecret == "" {
		return fmt.Errorf("client credentials not found")
	}

	var auth0Token auth.Auth0AuthToken

	token, err := auth0Token.GetOauthToken(*clientCreds)
	if err != nil {
		return fmt.Errorf("failed to get token: %v", err)
	}

	log.Println(token)

	return nil
}
