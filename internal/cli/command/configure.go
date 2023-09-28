package command

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/never00rei/go-auth0/internal/config"
	"github.com/urfave/cli/v2"
)

func CheckForEmptyInput(s string) error {
	if s == "" {
		return fmt.Errorf("input cannot be empty")
	}
	return nil
}

func CheckForWhitespace(s string) error {
	if strings.ContainsAny(s, " \t\r\n") {
		return fmt.Errorf("input cannot contain whitespace characters")
	}
	return nil
}

func tenantValidator(tenant string) error {
	if err := CheckForEmptyInput(tenant); err != nil {
		return err
	}

	if err := CheckForWhitespace(tenant); err != nil {
		return err
	}

	return nil
}

func clientIDValidator(input string) error {
	if err := CheckForEmptyInput(input); err != nil {
		return err
	}

	if err := CheckForWhitespace(input); err != nil {
		return err
	}

	return nil
}

func clientSecretValidator(input string) error {
	if err := CheckForEmptyInput(input); err != nil {
		return err
	}

	if err := CheckForWhitespace(input); err != nil {
		return err
	}

	if len(input) < 12 {
		return fmt.Errorf("input is too short, must be at least 12 characters long")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(input) {
		return fmt.Errorf("input should contain at least one uppercase letter")
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(input) {
		return fmt.Errorf("input should contain at least one lowercase letter")
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(input) {
		return fmt.Errorf("input should contain at least one number")
	}

	if !regexp.MustCompile(`[!@#$%^&*()]`).MatchString(input) { // Add more symbols as needed
		return fmt.Errorf("input should contain at least one special character")
	}

	return nil
}

func Configure(c *cli.Context) error {

	tenantPrompt := promptui.Prompt{
		Label:    "Auth0 Tenant",
		Validate: tenantValidator,
	}

	tenant, err := tenantPrompt.Run()
	if err != nil {
		return fmt.Errorf("failed to get tenant: %v", err)
	}

	clientIDPrompt := promptui.Prompt{
		Label:    "Auth0 Client ID",
		Validate: clientIDValidator,
	}

	clientID, err := clientIDPrompt.Run()
	if err != nil {
		return fmt.Errorf("failed to get client ID: %v", err)
	}

	clientSecretPrompt := promptui.Prompt{
		Label:    "Auth0 Client Secret",
		Validate: clientSecretValidator,
		Mask:     '*',
	}

	clientSecret, err := clientSecretPrompt.Run()
	if err != nil {
		return fmt.Errorf("failed to get client secret: %v", err)
	}

	clientApiServerPrompt := promptui.Prompt{
		Label: "Auth0 API Server",
	}

	clientApiServer, err := clientApiServerPrompt.Run()
	if err != nil {
		return fmt.Errorf("failed to get API server: %v", err)
	}

	cfg := &config.ClientAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Tenant:       tenant,
		ApiServer:    clientApiServer,
	}

	if err := config.SaveCredentialsFile(*cfg); err != nil {
		return fmt.Errorf("failed to save credentials file: %v", err)
	}

	return nil
}
