package dto

import "time"

type UserCreateData struct {
	UserStatusID int64
	Name         string
	Birthdate    time.Time
	Email        string
}
