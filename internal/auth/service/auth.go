package service

import (
	"context"
	"errors"
	"go-gin-ticketing-backend/internal/auth/domain"
	authrepository "go-gin-ticketing-backend/internal/auth/repository"
	"go-gin-ticketing-backend/internal/auth/schemas"
	"go-gin-ticketing-backend/internal/shared/enums"
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

func (s *AuthService) RegisterUser(ctx context.Context, body schemas.RegisterBody) (*domain.UserCredential, error) {

	birthdate, err := time.Parse("2006-01-02", body.Birthdate)
	if err != nil {
		return nil, err
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 12)

	user, err := domain.NewUserCredential(
		int64(enums.EmailConfirmation),
		body.Name,
		birthdate,
		body.Email,
		string(passwordHash),
	)
	if err != nil {
		return nil, err
	}

	return s.authRepository.RegisterUser(ctx, user)
}

func (s *AuthService) LoginUser(ctx context.Context, body schemas.LoginBody) (string, error) {

	user, err := s.authRepository.GetUserByEmail(ctx, body.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	switch user.UserInfo.UserStatusID {
	case int64(enums.Inactive):
		return "", errors.New("invalid credentials, inactive account")
	case int64(enums.PasswordCreation):
		return "", errors.New("invalid credentials, password creation pending")
	case int64(enums.EmailConfirmation):
		return "", errors.New("invalid credentials, email confirmation pending")
	case int64(enums.DeletedAccount):
		return "", errors.New("invalid credentials, deleted account")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)) != nil {
		return "", errors.New("invalid credentials")
	}

	claims := schemas.CustomClaims{
		Role: "system", // Change this later, setting up all users as role=system
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(user.UserInfo.ID, 10),
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
