package crypto_hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher(config Config) *BcryptHasher {
	return &BcryptHasher{cost: config.Cost}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", fmt.Errorf("process bcrypt: %w", err)
	}

	return string(bytes), nil
}

func (h *BcryptHasher) Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
