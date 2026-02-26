package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, handler *AuthHandler) {

	r.POST("/login", handler.LoginUser)
	r.POST("/register", handler.RegisterUser)
}
