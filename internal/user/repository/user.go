package repository

import (
	"context"

	shareddomain "go-gin-ticketing-backend/internal/shared/domain"
	"go-gin-ticketing-backend/internal/user/dto"
	"go-gin-ticketing-backend/internal/user/models"
	"go-gin-ticketing-backend/internal/user/schemas"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, pagination *shareddomain.Pagination) ([]models.User, *int64, error)
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	CreateUser(ctx context.Context, creationData *dto.CreationData) (*int64, error)
	UpdateUserByID(ctx context.Context, id int64, data schemas.UpdateUserBody) (*models.User, error)
	DeleteUserByID(ctx context.Context, id int64) (bool, error)
}
