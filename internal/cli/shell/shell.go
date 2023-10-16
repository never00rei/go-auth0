package shell

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	//"github.com/never00rei/go-auth0/internal/auth"
	"github.com/never00rei/go-auth0/internal/auth"
	"github.com/never00rei/go-auth0/internal/config"
)

type ShellEnvironment struct{}

func (s ShellEnvironment) GetDefaultShell() string {
	shell := os.Getenv("SHELL")
	return shell
}

// Not implemented yet...:
func (s ShellEnvironment) GetEnvironmentVars() error {

	return nil
}

func (s ShellEnvironment) NewSubShell(a auth.Auth0AuthToken) error {
	shell := s.GetDefaultShell()

	if shell == "" {
		return fmt.Errorf("failed to get default shell")
	}

	log.Println("Spawning new subshell for Go-auth0.")

	cmd := exec.Command(shell)

	env := os.Environ()

	tokenEnvVar := fmt.Sprintf("%s=%s", config.EnvSessionBearerToken, a.Token.OauthToken)
	expiryEnvVar := fmt.Sprintf("%s=%s", config.EnvSessionTokenExpiryTime, a.ExpiresDate.String())
	tenantEnvVar := fmt.Sprintf("%s=%s", config.EnvSessionTenant, a.ClientAuth.Tenant)

	env = append(env, tokenEnvVar)
	env = append(env, expiryEnvVar)
	env = append(env, tenantEnvVar)

	cmd.Env = env

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to start the shell: %s\n", err)
	}

	return nil

}
