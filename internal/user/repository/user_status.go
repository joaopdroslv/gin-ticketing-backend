package repository

import (
	"context"
	"go-gin-ticketing-backend/internal/user/models"
)

type UserStatusRepository interface {
	ListUserStatuses(ctx context.Context) ([]models.UserStatus, error)
}
