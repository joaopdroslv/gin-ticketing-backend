package dto

import "time"

type UserUpdateData struct {
	Name      *string
	Birthdate *time.Time
	Email     *string
}
