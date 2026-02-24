package models

type UserInfo struct {
	ID           int64
	UserStatusID int64
}

type UserCredential struct {
	ID           int64
	Email        string
	PasswordHash string

	UserInfo UserInfo
}
