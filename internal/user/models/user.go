package models

import "time"

type User struct {
	ID               int64
	UserCredentialID int64
	UserStatusID     int64
	Email            string
	Name             string
	Birthdate        time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
