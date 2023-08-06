package jwt

import "github.com/golang-jwt/jwt/v5"

type TokenUser struct {
	SteamID string `json:"steamid"`
	Admin   bool   `json:"admin"`
	jwt.RegisteredClaims
}
