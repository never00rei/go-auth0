package models

import "encoding/json"

type RefreshToken struct {
	Expiration_Type              string `json:"expiration_type"`
	Leeway                       int64  `json:"leeway"`
	Infinite_Token_Lifetime      bool   `json:"inifinite_token_lifetime"`
	Infinite_Idle_Token_Lifetime bool   `json:"infinite_idle_token_lifetime"`
	Token_Lifetime               int64  `json:"token_lifetime"`
	Idle_Token_Lifetime          int64  `json:"idle_token_lifetime"`
	Rotation_type                string `json:"rotation_type"`
}

type SigningKeys struct {
	Cert    string `json:"cert"`
	Pkcs7   string `json:"pkcs7"`
	Subject string `json:"subject"`
}

type JwtConfiguration struct {
	Alg                 string `json:"alg"`
	Lifetime_In_Seconds int64  `json:"lifetime_in_seconds"`
	Secret_Encoded      bool   `json:"secret_encoded"`
}

type Client struct {
	Tenant                              string           `json:"tenant"`
	Global                              bool             `json:"global"`
	Is_Token_Endpoint_Ip_Header_Trusted bool             `json:"is_token_endpoint_ip_header_trusted"`
	Name                                string           `json:"name"`
	Callbacks                           []any            `json:"callbacks"`
	Is_First_Party                      bool             `json:"is_first_party"`
	Oidc_Conformant                     bool             `json:"oidc_Conformant"`
	Sso_Disabled                        bool             `json:"sso_disabled"`
	Cross_Origin_Auth                   bool             `json:"cross_origin_auth"`
	Refresh_Token                       RefreshToken     `json:"refresh_token"`
	Signing_Keys                        []SigningKeys    `json:"signing_keys"`
	Owners                              []string         `json:"owners"`
	Client_Id                           string           `json:"client_id"`
	Callback_Url_Template               bool             `json:"callback_url_template"`
	Client_Secret                       string           `json:"client_secret"`
	Jwt_Configuration                   JwtConfiguration `json:"jwt_configuration"`
	Grant_Types                         []string         `json:"grant_types"`
	Custom_Login_Page_On                bool             `json:"custom_login_page_on"`
}

func (c *Client) UnmarshalToModel(data []byte) error {
	return json.Unmarshal(data, c)
}
