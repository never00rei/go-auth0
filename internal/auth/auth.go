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

func (b BearerToken) ConstructBearerToken() string {
	return fmt.Sprintf("%s %s", b.TokenType, b.OauthToken)
}

type Auth0AuthToken struct {
	ClientAuth  config.ClientAuth
	Token       BearerToken
	CreatedDate time.Time
	ExpiresDate time.Time
}

func GenerateTokenPayload(c config.ClientAuth, apiBaseUrl string) *strings.Reader {
	payload := strings.NewReader(
		fmt.Sprintf(
			`{"client_id":"%s","client_secret":"%s","audience":"%s","grant_type":"client_credentials"}`,
			c.ClientID,
			c.ClientSecret,
			apiBaseUrl,
		),
	)
	return payload
}

func GetOauthToken(c config.ClientAuth) (*Auth0AuthToken, error) {
	tokenUrl := config.Auth0ApiOauthUrl(c.ApiDomain)
	apiBaseUrl := config.Auth0ApiBaseUrl(c.ApiDomain)

	tokenPayload := GenerateTokenPayload(c, apiBaseUrl)

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
