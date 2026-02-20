package handler

import (
	"go-gin-ticketing-backend/internal/auth/schemas"
	"go-gin-ticketing-backend/internal/auth/service"
	sharedschemas "go-gin-ticketing-backend/internal/shared/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuthHandler struct {
	authService *service.AuthService
}

func New(authService *service.AuthService) *UserAuthHandler {

	return &UserAuthHandler{authService: authService}
}

func (h *UserAuthHandler) RegisterUser(c *gin.Context) {

	var body schemas.RegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		sharedschemas.Failed(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.authService.RegisterUser(c, body)
	if err != nil {
		sharedschemas.Failed(c, http.StatusInternalServerError, err.Error())
	}

	sharedschemas.OK(c, gin.H{"message": "user registered successfully"})
}

func (h *UserAuthHandler) LoginUser(c *gin.Context) {

	var body schemas.LoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		sharedschemas.Failed(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.authService.LoginUser(c, body)
	if err != nil {
		sharedschemas.Failed(c, http.StatusInternalServerError, err.Error())
		return
	}

	sharedschemas.OK(c, gin.H{"token": token})
}
