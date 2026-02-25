package service

import (
	"context"
	"go-gin-ticketing-backend/internal/access_control/models"
	"go-gin-ticketing-backend/internal/access_control/repository"
)

type PermissionService struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionService(
	permissionRepository repository.PermissionRepository,
) *PermissionService {

	return &PermissionService{permissionRepository: permissionRepository}
}

func (s *PermissionService) GetAllPermissions(
	ctx context.Context,
	name string,
) ([]models.Permission, error) {

	return s.permissionRepository.GetAllPermissions(ctx, name)
}

func (s *PermissionService) GetPermissionsByRoleID(
	ctx context.Context,
	id int64,
) ([]models.Permission, error) {

	return s.permissionRepository.GetPermissionsByRoleID(ctx, id)
}

func (s *PermissionService) UserHasPermission(
	ctx context.Context,
	userID int64,
	requiredPermission string,
) (bool, error) {

	return s.permissionRepository.UserHasPermission(ctx, userID, requiredPermission)
}
