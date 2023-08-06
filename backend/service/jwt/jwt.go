package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
)

type JWTService interface {
	IssueJWT(user *entity.User) (string, error)
	ValidateJWT(token string) (*TokenUser, error)
}

type jwtService struct {
	key []byte
}

func NewJWTGateway(key []byte) JWTService {
	return &jwtService{
		key: key,
	}
}

func (j *jwtService) IssueJWT(user *entity.User) (string, error) {
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

func (j *jwtService) ValidateJWT(token string) (*TokenUser, error) {
	claims := &TokenUser{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return j.key, nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
