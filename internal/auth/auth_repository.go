package auth

import (
	"context"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*UserCredential, error)
	RegisterUser(ctx context.Context, data *RegisterUserData) error
}
