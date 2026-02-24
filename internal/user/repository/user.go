package repository

import (
	"context"

	shareddomain "go-gin-ticketing-backend/internal/shared/domain"
	"go-gin-ticketing-backend/internal/user/dto"
	"go-gin-ticketing-backend/internal/user/models"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, pagination *shareddomain.Pagination) ([]models.User, *int64, error)
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	CreateUser(ctx context.Context, data *dto.UserCreateData) (*int64, error)
	UpdateUserByID(ctx context.Context, id int64, data *dto.UserUpdateData) (*models.User, error)
	DeleteUserByID(ctx context.Context, id int64) (bool, error)
}
