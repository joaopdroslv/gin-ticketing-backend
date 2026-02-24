package repository

import (
	"context"
	"go-gin-ticketing-backend/internal/auth/models"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.UserCredential, error)
	RegisterUser(ctx context.Context, user *models.RegistrationData) error
}
