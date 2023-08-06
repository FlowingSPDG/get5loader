package usecase

import "github.com/FlowingSPDG/get5-web-go/backend/service/jwt"

// ValidateJWT is interface for validating jwt token.
type ValidateJWT interface {
	Validate(token string) (isAdmin bool, err error)
}

type validateJWT struct {
	jwtService jwt.JWTService
}

func (vj *validateJWT) Validate(token string) (bool, error) {
	tokenUser, err := vj.jwtService.ValidateJWT(token)
	if err != nil {
		return false, err
	}

	return tokenUser.Admin, nil
}

func NewValidateJWT(
	jwtService jwt.JWTService,
) ValidateJWT {
	return &validateJWT{
		jwtService: jwtService,
	}
}
