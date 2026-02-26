package user

import (
	"context"

	"go-gin-ticketing-backend/internal/domain"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, pagination *domain.Pagination) ([]User, *int64, error)
	GetAllUserStatuses(ctx context.Context) ([]UserStatus, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, data *CreateUserData) (*int64, error)
	UpdateUserByID(ctx context.Context, id int64, data *UpdateUserData) (*User, error)
	DeleteUserByID(ctx context.Context, id int64) (bool, error)
}
