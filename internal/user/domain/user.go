package domain

import (
	"errors"
	"time"
)

type User struct {
	ID           int64
	UserStatusID int64
	Email        string
	Name         string
	Birthdate    time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(userStatusID int64, email, name string, birthdate time.Time) (*User, error) {

	if userStatusID <= 0 {
		return nil, errors.New("user_status_id is required")
	}
	if email == "" {
		return nil, errors.New("e-mail is required")
	}
	if name == "" {
		return nil, errors.New("name is required")
	}
	if birthdate.IsZero() {
		return nil, errors.New("birthdate is required")
	}

	return &User{
		UserStatusID: userStatusID,
		Email:        email,
		Name:         name,
		Birthdate:    birthdate,
	}, nil
}
