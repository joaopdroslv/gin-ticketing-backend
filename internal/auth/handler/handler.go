package handler

import (
	"go-gin-ticketing-backend/internal/auth/schemas"
	"go-gin-ticketing-backend/internal/auth/service"
	sharedschemas "go-gin-ticketing-backend/internal/shared/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuthHandler struct {
	service *service.UserAuthService
}

func New(service *service.UserAuthService) *UserAuthHandler {

	return &UserAuthHandler{service: service}
}

func (h *UserAuthHandler) RegisterUser(c *gin.Context) {

	var body schemas.UserRegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		sharedschemas.Failed(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.RegisterUser(c, body)
	if err != nil {
		sharedschemas.Failed(c, http.StatusInternalServerError, err.Error())
	}

	sharedschemas.OK(c, gin.H{"id": user.ID})
}

func (h *UserAuthHandler) LoginUser(c *gin.Context) {

	var body schemas.UserLoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		sharedschemas.Failed(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.LoginUser(c, body)
	if err != nil {
		sharedschemas.Failed(c, http.StatusInternalServerError, err.Error())
		return
	}

	sharedschemas.OK(c, gin.H{"token": token})
}
