package main

import (
	"fmt"
	"log"
	"os"

	"github.com/never00rei/go-auth0/internal/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  config.AppName,
		Usage: "A CLI tool for Auth0",
		Action: func(*cli.Context) error {
			fmt.Println("Hello world!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
