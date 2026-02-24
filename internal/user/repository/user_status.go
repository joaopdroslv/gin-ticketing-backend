package repository

import (
	"context"
	"go-gin-ticketing-backend/internal/user/models"
)

type UserStatusRepository interface {
	GetAllUserStatuses(ctx context.Context) ([]models.UserStatus, error)
}
