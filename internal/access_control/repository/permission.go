package repository

import (
	"context"
	"go-gin-ticketing-backend/internal/access_control/models"
)

type PermissionRepository interface {
	GetAllPermissions(ctx context.Context, name string) ([]models.Permission, error)
	GetPermissionsByRoleID(ctx context.Context, userID int64) ([]models.Permission, error)
	UserHasPermission(ctx context.Context, id int64, permission string) (bool, error)
}
