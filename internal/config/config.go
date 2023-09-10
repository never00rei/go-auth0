package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

const (
	ConfigFile   = "credentials"
	ConfigFolder = ".auth0"
	AppName      = "auth0-cli"
)

var HomeDir = os.Getenv("HOME")
var ConfigPath = fmt.Sprintf("%s/%s/%s", HomeDir, ConfigFolder, ConfigFile)

type ClientAuth struct {
	ClientID     string
	ClientSecret string
	Tenant       string
}

func CheckFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func ReadCredentialsFile(filePath string, tenant string) (*ini.File, error) {
	cfg, err := ini.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials file: %v", err)
	}

	return cfg, nil
}

func GetCredentials(tenant string) (*ClientAuth, error) {
	if !CheckFileExists(ConfigPath) {
		return nil, fmt.Errorf("credentials file not found")
	}

	cfg, err := ReadCredentialsFile(ConfigPath, tenant)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials file: %v", err)
	}

	section, err := cfg.GetSection(tenant)

	if err != nil {
		return nil, fmt.Errorf("failed to get credentials section: %v", err)
	}

	config := &ClientAuth{
		ClientID:     section.Key("ClientID").String(),
		ClientSecret: section.Key("ClientSecret").String(),
		Tenant:       tenant,
	}

	return config, nil
}

func SaveCredentialsFile(config ClientAuth) error {

	var credentials *ini.File
	var err error

	if !CheckFileExists(ConfigPath) {
		credentials = ini.Empty()
	} else {
		credentials, err = ReadCredentialsFile(ConfigPath, config.Tenant)
		if err != nil {
			return fmt.Errorf("failed to read credentials file: %v", err)
		}
	}

	section, err := credentials.NewSection(config.Tenant)
	if err != nil {
		return fmt.Errorf("failed to create new section: %v", err)
	}

	if _, err = section.NewKey("ClientID", config.ClientID); err != nil {
		return fmt.Errorf("failed to create new key: %v", err)
	}

	if _, err = section.NewKey("ClientSecret", config.ClientSecret); err != nil {
		return fmt.Errorf("failed to create new key: %v", err)
	}

	if err = credentials.SaveTo(ConfigPath); err != nil {
		return fmt.Errorf("failed to save INI data: %v", err)
	}

	return nil
}
