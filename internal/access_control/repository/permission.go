package repository

import (
	"context"
	"go-gin-ticketing-backend/internal/access_control/models"
)

type PermissionRepository interface {
	GetPermissionsByUserID(ctx context.Context, userID int64) ([]models.Permission, error)
}
