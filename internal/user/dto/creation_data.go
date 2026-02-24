package dto

import "time"

type CreationData struct {
	UserStatusID int64
	Name         string
	Birthdate    time.Time
	Email        string
}
