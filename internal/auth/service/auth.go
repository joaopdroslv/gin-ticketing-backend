package service

import (
	"context"
	"go-gin-ticketing-backend/internal/auth/models"
	authrepository "go-gin-ticketing-backend/internal/auth/repository"
	"go-gin-ticketing-backend/internal/auth/schemas"
	"go-gin-ticketing-backend/internal/shared/enums"
	"go-gin-ticketing-backend/internal/shared/errs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository       authrepository.AuthRepository
	permissionRepository authrepository.PermissionRepository
	jwtSecret            []byte
	jwtTTL               time.Duration
}

func New(
	authRepository authrepository.AuthRepository,
	permissionRepository authrepository.PermissionRepository,
	jwtSecret string,
	jwtTTL int64,
) *AuthService {

	return &AuthService{
		authRepository:       authRepository,
		permissionRepository: permissionRepository,
		jwtSecret:            []byte(jwtSecret),
		jwtTTL:               time.Duration(jwtTTL) * time.Second,
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, body schemas.RegisterBody) error {

	birthdate, err := time.Parse("2006-01-02", body.Birthdate)
	if err != nil {
		return err
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 12)

	registrationData := &models.RegistrationData{
		UserStatusID: int64(enums.Active),
		Name:         body.Name,
		Birthdate:    birthdate,
		Email:        body.Email,
		PasswordHash: string(passwordHash),
	}

	return s.authRepository.RegisterUser(ctx, registrationData)
}

func (s *AuthService) LoginUser(ctx context.Context, body schemas.LoginBody) (string, error) {

	userCredential, err := s.authRepository.GetUserByEmail(ctx, body.Email)
	if err != nil {
		return "", errs.ErrInvalidCredentials
	}

	err = s.validateUserStatus(userCredential)
	if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(userCredential.PasswordHash), []byte(body.Password)) != nil {
		return "", errs.ErrInvalidCredentials
	}

	claims := schemas.CustomClaims{
		Role: "system", // Change this later, setting up all users as role=system
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(userCredential.UserInfo.ID, 10),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.jwtTTL)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) HasThisPermission(ctx context.Context, userID int64, userPermission string) (bool, error) {

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
	_, ok := permissionsMap[userPermission]

	return ok, nil
}

func (s *AuthService) validateUserStatus(userCredential *models.UserCredential) error {

	switch userCredential.UserInfo.UserStatusID {
	case int64(enums.Inactive):
		return errs.ErrInactiveUser
	case int64(enums.EmailConfirmationPending):
		return errs.ErrUserEmailConfirmationPending
	case int64(enums.PasswordCreationPending):
		return errs.ErrUserPasswordCreationPending
	case int64(enums.Deleted):
		return errs.ErrDeletedUser
	}

	return nil
}
