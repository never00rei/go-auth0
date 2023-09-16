package cli

import (
	"log"
	"os"

	"github.com/never00rei/go-auth0/internal/cli/command"
	"github.com/never00rei/go-auth0/internal/config"
	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:  config.AppName,
		Usage: "A CLI tool for Auth0",
		Commands: []*cli.Command{
			{
				Name:   "configure",
				Usage:  "Configure the CLI tool",
				Action: command.Configure,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
