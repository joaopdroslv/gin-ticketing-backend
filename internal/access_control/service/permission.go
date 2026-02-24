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

	// Step 1. Get all user's userPermissions using its ID
	userPermissions, err := s.permissionRepository.GetPermissionsByUserID(ctx, userID)
	if err != nil {
		return false, err
	}

	// Step 2. Creating a permissions map (with empty structs) for each user permissions
	permissionsMap := make(map[string]struct{})

	for _, permission := range userPermissions {
		permissionsMap[permission.Name] = struct{}{}
	}

	// Step 3. Validating if the user has the required permission
	_, ok := permissionsMap[requiredPermission]

	return ok, nil
}
