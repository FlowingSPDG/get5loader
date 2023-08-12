package jwt_test

import (
	"testing"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/service/jwt"
	"github.com/stretchr/testify/assert"
)

func TestIssueJWT(t *testing.T) {
	tt := []struct {
		name     string
		input    *entity.User
		expected *entity.TokenUser
	}{
		{
			name: "success",
			input: &entity.User{
				ID:      "1",
				SteamID: 76561198072054549,
				Admin:   true,
			},
			expected: &entity.TokenUser{
				UserID:  "1",
				SteamID: 76561198072054549,
				Admin:   true,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			j := jwt.NewJWTGateway([]byte("test"))
			actual, err := j.IssueJWT(tc.input)
			assert.NoError(t, err)

			token, err := j.ValidateJWT(actual)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, token)
		})
	}
}
