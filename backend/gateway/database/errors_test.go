package database_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

func TestErrors(t *testing.T) {
	tc := []struct {
		name     string
		input    error
		method   func(error) bool
		expected bool
	}{
		// default errors
		{
			name:     "ErrNotFound",
			method:   database.IsNotFound,
			input:    database.ErrNotFound,
			expected: true,
		},
		{
			name:     "ErrInteral",
			method:   database.IsInternal,
			input:    database.ErrInternal,
			expected: true,
		},

		// wrapped errors
		{
			name:     "wrapped ErrNotFound",
			method:   database.IsNotFound,
			input:    fmt.Errorf("User not found %w", database.ErrNotFound),
			expected: true,
		},
		{
			name:     "wrapped ErrInternal",
			method:   database.IsInternal,
			input:    fmt.Errorf("User not found %w", database.ErrInternal),
			expected: true,
		},

		// wrapped wrapped errors
		{
			name:     "wrapped wrapped ErrNotFound",
			method:   database.IsNotFound,
			input:    fmt.Errorf("User not found %w", fmt.Errorf("User not found %w", database.ErrNotFound)),
			expected: true,
		},
		{
			name:     "wrapped wrapped ErrInternal",
			method:   database.IsInternal,
			input:    fmt.Errorf("User not found %w", fmt.Errorf("User not found %w", database.ErrInternal)),
			expected: true,
		},

		// Different errors
		{
			name:     "different error",
			method:   database.IsNotFound,
			input:    fmt.Errorf("Unknown error!"),
			expected: false,
		},

		// wrapped different errors
		{
			name:     "wrapped different error",
			method:   database.IsNotFound,
			input:    database.ErrInternal,
			expected: false,
		},
		{
			name:     "wrapped different error",
			method:   database.IsInternal,
			input:    database.ErrNotFound,
			expected: false,
		},
	}

	for _, c := range tc {
		assert.Equal(t, c.expected, c.method(c.input), c.name)
	}
}
