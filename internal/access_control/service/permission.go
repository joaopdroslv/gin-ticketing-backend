package service

import (
	"context"
	accessrepository "go-gin-ticketing-backend/internal/access_control/repository"
)

type PermissionService struct {
	permissionRepository accessrepository.PermissionRepository
}

func NewPermissionService(permissionRepository accessrepository.PermissionRepository) *PermissionService {

	return &PermissionService{permissionRepository: permissionRepository}
}

func (s *PermissionService) UserHasPermission(
	ctx context.Context,
	userID int64,
	requiredPermission string,
) (bool, error) {

	return s.permissionRepository.UserHasPermission(ctx, userID, requiredPermission)
}
