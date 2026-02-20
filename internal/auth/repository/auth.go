package repository

import (
	"context"
	"go-gin-ticketing-backend/internal/auth/domain"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.UserCredential, error)
	RegisterUser(ctx context.Context, user *domain.UserCredential) (*domain.UserCredential, error)
}
