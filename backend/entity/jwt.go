package entity

import "github.com/golang-jwt/jwt/v5"

type TokenUser struct {
	UserID  UserID  `json:"userid"`
	SteamID SteamID `json:"steamid"`
	Admin   bool    `json:"admin"`
	jwt.RegisteredClaims
}
