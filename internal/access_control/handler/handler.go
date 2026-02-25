package handler

import "go-gin-ticketing-backend/internal/access_control/service"

type PermissionHandler struct {
	permissionService service.PermissionService
}

func New(permissionService service.PermissionService) *PermissionHandler {

	return &PermissionHandler{permissionService: permissionService}
}
