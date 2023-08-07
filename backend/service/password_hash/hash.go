package hash

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	Hash(password string) ([]byte, error)
	Compare(hash []byte, password string) error
}

type passwordHasher struct {
	cost int
}

func NewPasswordHasher(cost int) PasswordHasher {
	return &passwordHasher{
		cost: cost,
	}
}

// Compare implements PasswordHasher.
func (ph *passwordHasher) Compare(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

// Hash implements PasswordHasher.
func (ph *passwordHasher) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), ph.cost)
}
