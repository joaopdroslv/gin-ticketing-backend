package models

import "time"

type User struct {
	ID               int64
	UserCredentialID int64
	UserStatusID     int64
	Name             string
	Birthdate        time.Time
	Email            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
