package domain

import (
	"errors"
	"time"
)

// TODO: Think about this, maybe its not needed
type UserInfo struct {
	ID           int64
	UserStatusID int64
	Name         string
	Birthdate    time.Time
}

type UserCredential struct {
	ID           int64
	Email        string
	PasswordHash string
	UserInfo     UserInfo
}

func NewUserCredential(userStatusID int64, name string, birthdate time.Time, email, passwordHash string) (*UserCredential, error) {

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

	return &UserCredential{
		Email:        email,
		PasswordHash: passwordHash,
		UserInfo: UserInfo{
			UserStatusID: userStatusID,
			Name:         name,
			Birthdate:    birthdate,
		},
	}, nil
}
