package repository

import (
	"context"
	"ticket-io/internal/auth/domain"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*domain.AuthUser, error)
}
