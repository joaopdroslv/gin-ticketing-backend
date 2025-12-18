package domain

import "errors"

type AuthUser struct {
	ID           int64
	Email        string
	PasswordHash string
}

func NewAuthUser(email, passwordHash string) (*AuthUser, error) {

	if email == "" {
		return nil, errors.New("e-mail is required")
	}
	if passwordHash == "" {
		return nil, errors.New("passwordHash is required")
	}

	return &AuthUser{
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}
