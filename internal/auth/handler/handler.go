package handler

import (
	"errors"
	"go-gin-ticketing-backend/internal/auth/schemas"
	"go-gin-ticketing-backend/internal/auth/service"
	"go-gin-ticketing-backend/internal/shared/errs"
	sharedschemas "go-gin-ticketing-backend/internal/shared/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func New(authService *service.AuthService) *AuthHandler {

	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {

	var body schemas.RegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		sharedschemas.Failed(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.authService.RegisterUser(c, body)
	if err != nil {
		sharedschemas.Failed(c, http.StatusInternalServerError, err.Error())
	}

	sharedschemas.OK(c, gin.H{"message": "user registered successfully"})
}

func (h *AuthHandler) LoginUser(c *gin.Context) {

	var body schemas.LoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		sharedschemas.Failed(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.authService.LoginUser(c, body)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidCredentials) {
			sharedschemas.Failed(c, http.StatusUnauthorized, err.Error())
			return
		}

		if errs.IsUserStatusRelated(err) {
			sharedschemas.Failed(c, http.StatusForbidden, err.Error())
			return
		}

		sharedschemas.Failed(c, http.StatusInternalServerError, err.Error())
		return
	}

	sharedschemas.OK(c, gin.H{"message": "logged in successfully", "token": token})
}
