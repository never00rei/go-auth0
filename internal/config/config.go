package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

const (
	ConfigFile                string = "credentials"
	ConfigFolder              string = ".auth0"
	AppName                   string = "auth0-cli"
	EnvSessionBearerToken     string = "AUTH0_SESSION_TOKEN"
	EnvSessionTokenExpiryTime string = "AUTH0_SESSION_EXPIRY"
	EnvSessionTenant          string = "AUTH0_TENANT"
	EnvSessionApiUrl          string = "AUTH0_SESSION_API_URL"
	Auth0ApiVersion           string = "v2"
)

var HomeDir = os.Getenv("HOME")
var ConfigPath = fmt.Sprintf("%s/%s/%s", HomeDir, ConfigFolder, ConfigFile)

type ClientAuth struct {
	ClientID     string
	ClientSecret string
	Tenant       string
	ApiDomain    string
}

type FileSystem interface {
	Stat(name string) (os.FileInfo, error)
}

type OSFileSystem struct{}

func (OSFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func CheckFileExists(fs FileSystem, filePath string) bool {
	if _, err := fs.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateConfigFolder() {
	err := os.MkdirAll(fmt.Sprintf("%s/%s", HomeDir, ConfigFolder), 0700)
	if err != nil {
		log.Print("Error creating directory")
	}
}

func ReadCredentialsFile(filePath string, tenant string) (*ini.File, error) {
	cfg, err := ini.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials file: %v", err)
	}

	return cfg, nil
}

func GetCredentials(tenant string) (*ClientAuth, error) {

	fs := OSFileSystem{}

	if !CheckFileExists(fs, ConfigPath) {
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
		Tenant:       section.Name(),
		ApiDomain:    section.Key("ApiDomain").String(),
	}

	return config, nil
}

func SaveCredentialsFile(config ClientAuth) error {

	var credentials *ini.File
	var err error

	fs := OSFileSystem{}

	if !CheckFileExists(fs, ConfigPath) {
		CreateConfigFolder()
		credentials = ini.Empty()
	} else {
		credentials, err = ReadCredentialsFile(ConfigPath, config.Tenant)
		if err != nil {
			return fmt.Errorf("failed to read credentials file: %v", err)
		}
	}

	// New section also updates an existing section if found.
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

	if _, err = section.NewKey("ApiDomain", config.ApiDomain); err != nil {
		return fmt.Errorf("failed to create new : %v", err)
	}

	if err = credentials.SaveTo(ConfigPath); err != nil {
		return fmt.Errorf("failed to save INI data: %v", err)
	}

	return nil
}

func Auth0ApiOauthUrl(domain string) string {
	return fmt.Sprintf("https://%s/oauth/token", domain)
}

func Auth0ApiBaseUrl(domain string) string {
	return fmt.Sprintf("https://%s/api/%s/", domain, Auth0ApiVersion)
}
