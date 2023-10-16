package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/never00rei/go-auth0/internal/config"
)

type BearerToken struct {
	OauthToken string `json:"access_token"`
	TokenType  string `json:"token_type"`
}

type Auth0AuthToken struct {
	ClientAuth  config.ClientAuth
	Token       BearerToken
	CreatedDate time.Time
	ExpiresDate time.Time
}

func GetOauthToken(c config.ClientAuth) (*Auth0AuthToken, error) {
	tokenUrl := fmt.Sprintf("https://%s/oauth/token", c.ApiDomain)

	tokenPayload := strings.NewReader(
		fmt.Sprintf(
			`{"client_id":"%s","client_secret":"%s","audience":"https://%s/api/v2/","grant_type":"client_credentials"}`,
			c.ClientID,
			c.ClientSecret,
			c.ApiDomain,
		),
	)

	req, err := http.NewRequest("POST", tokenUrl, tokenPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get token: %v", res.Status)
	}

	var bearerToken BearerToken

	if err := json.NewDecoder(res.Body).Decode(&bearerToken); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	a := &Auth0AuthToken{}

	a.ClientAuth = c
	a.CreatedDate = time.Now()
	a.ExpiresDate = a.CreatedDate.Add(time.Second * 36000)
	a.Token = bearerToken

	return a, nil
}

func (a Auth0AuthToken) IsTokenExpired() bool {
	return time.Now().After(a.ExpiresDate)
}
