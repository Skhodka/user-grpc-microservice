package domain

import "errors"

var (
	ErrEmptyHashPassword = errors.New("empty hash password")
)

type HashPassword struct {
	hash_pass string
}

func NewHashPassword(hash_pass string) (*HashPassword, error) {
	if hash_pass == "" {
		return nil, ErrEmptyHashPassword
	}

	return &HashPassword{hash_pass: hash_pass}, nil
}

func (h *HashPassword) GetPass() string {
	return h.hash_pass
}
