package repository

import (
	"context"

	"go-gin-ticketing-backend/internal/domain"
	"go-gin-ticketing-backend/internal/user/dto"
	"go-gin-ticketing-backend/internal/user/models"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, pagination *domain.Pagination) ([]models.User, *int64, error)
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	CreateUser(ctx context.Context, data *dto.CreateUserData) (*int64, error)
	UpdateUserByID(ctx context.Context, id int64, data *dto.UpdateUserData) (*models.User, error)
	DeleteUserByID(ctx context.Context, id int64) (bool, error)
}
