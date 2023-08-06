package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
)

type JWTGateway interface {
	IssueJWT(user *entity.User) (string, error)
	ValidateJWT(token string) (*TokenUser, error)
}

type jwtGateway struct {
	key []byte
}

func NewJWTGateway(key []byte) JWTGateway {
	return &jwtGateway{
		key: key,
	}
}

func (j *jwtGateway) IssueJWT(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &TokenUser{
		SteamID: user.SteamID,
		Admin:   user.Admin,
	})

	signed, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (j *jwtGateway) ValidateJWT(token string) (*TokenUser, error) {
	claims := &TokenUser{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
