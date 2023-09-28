package auth

import (
	"time"

	"github.com/never00rei/go-auth0/internal/config"
)

type Auth0AuthToken struct {
	ClientAuth         config.ClientAuth
	OauthToken         string
	CreatedDate        time.Time
	ExpiresDate        time.Time
	RefreshToken       string
	RefreshExpiresDate time.Time
}

func GetOauthToken(c config.ClientAuth) (string, error) {
	return "", nil
}
