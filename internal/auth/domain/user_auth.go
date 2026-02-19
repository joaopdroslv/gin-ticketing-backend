package domain

import (
	"errors"
	"time"
)

type UserAuth struct {
	ID           int64
	UserStatusID int64
	Name         string
	Birthdate    time.Time
	Email        string
	PasswordHash string
}

func NewUserAuth(userStatusID int64, name string, birthdate time.Time, email, passwordHash string) (*UserAuth, error) {

	if userStatusID <= 0 {
		return nil, errors.New("user_status_id is required")
	}
	if name == "" {
		return nil, errors.New("name is required")
	}
	if birthdate.IsZero() {
		return nil, errors.New("birthdate is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}
	if passwordHash == "" {
		return nil, errors.New("password_hash is required")
	}

	return &UserAuth{
		Name:         name,
		Birthdate:    birthdate,
		UserStatusID: userStatusID,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}
