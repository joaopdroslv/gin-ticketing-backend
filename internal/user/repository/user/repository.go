package user

import (
	"context"

	"ticket-io/internal/user/domain"
	"ticket-io/internal/user/dto"
)

type UserRepository interface {
	ListUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUserByID(ctx context.Context, id int64, data dto.UserUpdateBody) (*domain.User, error)
	DeleteUserByID(ctx context.Context, id int64) (bool, error)
}
