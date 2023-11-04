package models

import "encoding/json"

//[
//  {
//    "id": "con_V524HBuOXRBVWf2i",
//    "options": {
//      "email": true,
//      "scope": [
//        "email",
//        "profile"
//      ],
//      "profile": true
//    },
//    "strategy": "google-oauth2",
//    "name": "google-oauth2",
//    "is_domain_connection": false,
//    "realms": [
//      "google-oauth2"
//    ],
//    "enabled_clients": [
//      "dV5FrNRgWSxxbCRYo7jYkTHbhtCuO9Qj",
//      "6OF1j5qBdjXmc8FQSEw9z5MuWkcS6IKC"
//    ]
//  },
//  {
//    "id": "con_qvdECd0SGe21LdHz",
//    "options": {
//      "mfa": {
//        "active": true,
//        "return_enroll_settings": true
//      },
//      "passwordPolicy": "good",
//      "passkey_options": {
//        "challenge_ui": "both",
//        "local_enrollment_enabled": true,
//        "progressive_enrollment_enabled": true
//      },
//      "strategy_version": 2,
//      "authentication_methods": {
//        "passkey": {
//          "enabled": false
//       },
//        "password": {
//         "enabled": true
//        }
//      },
//      "brute_force_protection": true
//    },
//    "strategy": "auth0",
//    "name": "Username-Password-Authentication",
//    "is_domain_connection": false,
//   "realms": [
//      "Username-Password-Authentication"
//    ],
//    "enabled_clients": [
//      "dV5FrNRgWSxxbCRYo7jYkTHbhtCuO9Qj",
//      "6OF1j5qBdjXmc8FQSEw9z5MuWkcS6IKC"
//    ]
//  }
//]

// I don't like the fact that 'Options' has a type of 'any',
// other than spending hours trying to _find_ all of the options -
// I don't see any other way to do this...
type Connection struct {
	Id                   string         `json:"id"`
	Options              map[string]any `json:"options"`
	Strategy             string         `json:"strategy"`
	Name                 string         `json:"name"`
	Is_Domain_Connection bool           `json:"is_domain_connection"`
	Realms               []string       `json:"realms"`
	Enabled_Clients      []string       `json:"enabled_clients"`
}

func (c *Connection) UnmarshalToModel(data []byte) error {
	return json.Unmarshal(data, c)
}
