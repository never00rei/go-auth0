package models

import (
	"fmt"
	"time"
)

// Example response body
//[
//  {
//    "created_at": "2023-10-17T21:09:19.428Z",
//    "email": "jwashington104.jw+1@gmail.com",
//    "email_verified": false,
//    "identities": [
//      {
//        "connection": "Username-Password-Authentication",
//        "user_id": "652ef7ff158f5fdf275aa15a",
//       "provider": "auth0",
//        "isSocial": false
//      }
//    ],
//   "name": "jwashington104.jw+1@gmail.com",
//    "nickname": "jwashington104.jw+1",
//    "picture": "https://s.gravatar.com/avatar/9f826bf82d11d4b29aa337c2bf5b3191?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fjw.png",
//    "updated_at": "2023-10-17T21:09:19.428Z",
//    "user_id": "auth0|652ef7ff158f5fdf275aa15a"
//  }
//]

type Identity struct {
	Connection string
	User_Id    string
	Provider   string
	IsSocial   bool
}

func (i Identity) GetUserId() string {
	return fmt.Sprintf("%s|%s", i.Provider, i.User_Id)
}

type User struct {
	Created_At     time.Time
	Email          string
	Email_Verified bool
	Identities     []Identity
	Name           string
	Nickname       string
	Updated_at     time.Time
	Picture        string
	User_Id        string
}

func (u User) GetConnectionByUser() *string {

	return nil
}
