package dto

import "time"

type CreateUserData struct {
	UserStatusID int64
	Name         string
	Birthdate    time.Time
	Email        string
}

type UpdateUserData struct {
	Name      *string
	Birthdate *time.Time
	Email     *string
}
