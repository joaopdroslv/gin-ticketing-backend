package auth

import "time"

// Complementary info
type UserInfo struct {
	ID           int64
	UserStatusID int64
}

type UserCredential struct {
	ID           int64
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	UserInfo UserInfo
}
