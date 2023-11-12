package cli

import (
	"log"
	"os"

	"github.com/never00rei/go-auth0/internal/cli/command"
	"github.com/never00rei/go-auth0/internal/cli/command/clients"
	"github.com/never00rei/go-auth0/internal/cli/command/connections"
	"github.com/never00rei/go-auth0/internal/cli/command/users"
	"github.com/never00rei/go-auth0/internal/config"
	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:  config.AppName,
		Usage: "A CLI tool for Auth0",
		Commands: []*cli.Command{
			{
				Name:     "configure",
				Usage:    "Configure the CLI tool",
				Category: "Setup",
				Action:   command.Configure,
			},
			{
				Name:     "login",
				Usage:    "Login to Auth0",
				Category: "Setup",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "tenant",
						Aliases:  []string{"t"},
						Usage:    "Auth0 tenant",
						Required: true,
						EnvVars:  []string{config.EnvSessionTenant},
					},
				},
				Action: command.Login,
			},
			{
				Name:     "get-users",
				Category: "User Details",
				Usage:    "Fetches user details from Auth0 Management API.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "tenant",
						Aliases:  []string{"t"},
						Usage:    "Auth0 tenant",
						Required: true,
						EnvVars:  []string{config.EnvSessionTenant},
					},
				},
				Action: users.GetAllUsers,
			},
			{
				Name:     "get-connections",
				Category: "Connection Details",
				Usage:    "Fetches connection details from Auth0 Management API.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "tenant",
						Aliases:  []string{"t"},
						Usage:    "Auth0 tenant",
						Required: true,
						EnvVars:  []string{config.EnvSessionTenant},
					},
				},
				Action: connections.GetAllConnections,
			},
			{
				Name:     "get-clients",
				Category: "Client Details",
				Usage:    "Fetches client details from Auth0 Management API.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "tenant",
						Aliases:  []string{"t"},
						Usage:    "Auth0 tenant",
						Required: true,
						EnvVars:  []string{config.EnvSessionTenant},
					},
				},
				Action: clients.GetAllClients,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
