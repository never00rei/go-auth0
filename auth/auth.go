package auth

var AuthMethods = [2]string{"token", "client"}

// Here we define the auth methods that we can use to authenticate to the
// 	Auth0 Management API.

type TokenAuthMethod struct {
	Token string
}

type ClientAuthMethod struct {
	ClientID     string
	ClientSecret string
}

type Auth struct {
	Tenant     string
	AuthMethod string
	AuthData   map[string]string
}

type AuthData struct {
	Token  TokenAuthMethod
	Client ClientAuthMethod
}

func NewAuth(tenant, string, authMethod string, authData map[string]string) *Auth {
	return &Auth{
		Tenant:     tenant,
		AuthMethod: authMethod,
		AuthData:   authData,
	}
}
