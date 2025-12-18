package handler

import (
	"github.com/gin-gonic/gin"

	"ticket-io/internal/user/service/user"
)

func RegisterRoutes(r *gin.RouterGroup, userService *user.UserService) {

	userHandler := New(userService)

	r.GET("/users", userHandler.ListUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)
	r.POST("/users/:id", userHandler.UpdateUserByID)
	r.DELETE("/users/:id", userHandler.DeleteUserByID)
}
