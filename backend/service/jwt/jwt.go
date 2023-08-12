package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/FlowingSPDG/get5loader/backend/entity"
)

type JWTService interface {
	IssueJWT(userID entity.UserID, steamID entity.SteamID, admin bool) (string, error)
	ValidateJWT(token string) (*entity.TokenUser, error)
}

type jwtService struct {
	key []byte
}

func NewJWTGateway(key []byte) JWTService {
	return &jwtService{
		key: key,
	}
}

func (j *jwtService) IssueJWT(userID entity.UserID, steamID entity.SteamID, admin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &entity.TokenUser{
		UserID:  userID,
		SteamID: steamID,
		Admin:   admin,
	})

	signed, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (j *jwtService) ValidateJWT(token string) (*entity.TokenUser, error) {
	claims := &entity.TokenUser{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return j.key, nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
