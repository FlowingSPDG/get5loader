package usecase

import (
	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/service/jwt"
)

// ValidateJWT is interface for validating jwt token.
type ValidateJWT interface {
	Validate(token string) (*entity.TokenUser, error)
}

type validateJWT struct {
	jwtService jwt.JWTService
}

func (vj *validateJWT) Validate(token string) (*entity.TokenUser, error) {
	tokenUser, err := vj.jwtService.ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	return tokenUser, nil
}

func NewValidateJWT(
	jwtService jwt.JWTService,
) ValidateJWT {
	return &validateJWT{
		jwtService: jwtService,
	}
}
