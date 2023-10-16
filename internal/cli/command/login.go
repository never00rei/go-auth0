package command

import (
	"fmt"

	"github.com/never00rei/go-auth0/internal/auth"
	"github.com/never00rei/go-auth0/internal/cli/shell"
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

	authtoken, err := auth.GetOauthToken(*clientCreds)
	if err != nil {
		return fmt.Errorf("failed to get token: %v", err)
	}

	var sh shell.ShellEnvironment

	sherr := sh.NewSubShell(*authtoken)
	if sherr != nil {
		return fmt.Errorf("Failed to create shell: %v", sherr)
	}
	return nil
}
