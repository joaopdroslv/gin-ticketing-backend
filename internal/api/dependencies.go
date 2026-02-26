package api

import (
	accesscontrolservice "go-gin-ticketing-backend/internal/access_control/service"
	"go-gin-ticketing-backend/internal/auth"
	"go-gin-ticketing-backend/internal/user"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	AuthHandler       *auth.AuthHandler
	UserHandler       *user.UserHandler
	JWTMiddleware     *gin.HandlerFunc
	PermissionService *accesscontrolservice.PermissionService
}
