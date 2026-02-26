package api

import (
	"go-gin-ticketing-backend/internal/auth"
	"go-gin-ticketing-backend/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterV1(apiGroup *gin.RouterGroup, dependencies Dependencies) {

	v1Group := apiGroup.Group("/v1")

	authGroup := v1Group.Group("/auth")
	auth.RegisterAuthRoutes(authGroup, dependencies.AuthHandler)

	userGroup := v1Group.Group("/users")
	userGroup.Use(*dependencies.JWTMiddleware)
	user.RegisterUserRoutes(userGroup, dependencies.UserHandler, dependencies.PermissionService)
}
