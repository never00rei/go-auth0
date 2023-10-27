package models

type Model interface{}

var ModelRegistry = map[string]Model{
	"/users": &UserDetails{},
}
