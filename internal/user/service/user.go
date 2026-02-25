package service

import (
	"context"
	shareddomain "go-gin-ticketing-backend/internal/shared/domain"
	"go-gin-ticketing-backend/internal/shared/enums"
	sharedschemas "go-gin-ticketing-backend/internal/shared/schemas"
	"go-gin-ticketing-backend/internal/user/dto"
	"go-gin-ticketing-backend/internal/user/models"
	userrepository "go-gin-ticketing-backend/internal/user/repository"
	"go-gin-ticketing-backend/internal/user/schemas"
	"time"
)

type UserStatusProvider interface {
	GetUserStatusesMap(ctx context.Context) (map[int64]string, error)
}

type UserService struct {
	userRepository  userrepository.UserRepository
	userStatusesMap map[int64]string
}

func NewUserService(
	ctx context.Context,
	userRepository userrepository.UserRepository,
	userStatusProvider UserStatusProvider,
) (*UserService, error) {

	userStatusesMap, err := userStatusProvider.GetUserStatusesMap(ctx)
	if err != nil {
		return nil, err
	}

	return &UserService{
		userRepository:  userRepository,
		userStatusesMap: userStatusesMap,
	}, nil
}

func (s *UserService) GetAllUsers(
	ctx context.Context,
	paginationQuery sharedschemas.PaginationQuery,
) (*schemas.GetAllUsersResponse, error) {

	pagination := shareddomain.NewPagination(paginationQuery.Page, paginationQuery.Limit)

	users, total, err := s.userRepository.GetAllUsers(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return &schemas.GetAllUsersResponse{
		Items: s.translateUsers(users, s.userStatusesMap),
		Pagination: sharedschemas.ResponsePagination{
			Page:      pagination.Page,
			PageTotal: int64(len(users)),
			Limit:     pagination.Limit,
			Total:     *total,
		},
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*schemas.ResponseUser, error) {

	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.translateUser(user, s.userStatusesMap), nil
}

func (s *UserService) CreateUser(
	ctx context.Context,
	body schemas.CreateUserBody,
) (*schemas.ResponseUser, error) {

	birthdate, err := time.Parse("2006-01-02", body.Birthdate)
	if err != nil {
		return nil, err
	}

	data := &dto.CreateUserData{
		UserStatusID: int64(enums.PasswordCreationPending),
		Name:         body.Name,
		Birthdate:    birthdate,
		Email:        body.Email,
	}

	id, err := s.userRepository.CreateUser(ctx, data)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.GetUserByID(ctx, *id)
	if err != nil {
		return nil, err
	}

	return s.translateUser(user, s.userStatusesMap), nil
}

func (s *UserService) UpdateUserByID(
	ctx context.Context,
	id int64,
	body schemas.UpdateUserBody,
) (*schemas.ResponseUser, error) {

	data := &dto.UpdateUserData{}

	if body.Name != nil {
		data.Name = body.Name
	}
	if body.Birthdate != nil {
		birthdate, err := time.Parse("2006-01-02", *body.Birthdate)
		if err != nil {
			return nil, err
		}
		data.Birthdate = &birthdate
	}
	if body.Email != nil {
		data.Email = body.Email
	}

	user, err := s.userRepository.UpdateUserByID(ctx, id, data)
	if err != nil {
		return nil, err
	}

	return s.translateUser(user, s.userStatusesMap), nil
}

func (s *UserService) DeleteUserByID(ctx context.Context, id int64) (*schemas.DeleteUserResponse, error) {

	success, err := s.userRepository.DeleteUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &schemas.DeleteUserResponse{
		ID:      id,
		Deleted: success,
	}, nil
}

func (s *UserService) transformUserModelIntoResponseUser(
	userModel *models.User,
	userStatusName string,
) *schemas.ResponseUser {

	return &schemas.ResponseUser{
		ID:         userModel.ID,
		Name:       userModel.Name,
		Birthdate:  userModel.Birthdate.Format(time.RFC3339),
		Email:      userModel.Email,
		UserStatus: userStatusName,
	}
}

func (s *UserService) translateUser(
	userModel *models.User,
	userStatusesMap map[int64]string,
) *schemas.ResponseUser {

	return s.transformUserModelIntoResponseUser(userModel, userStatusesMap[userModel.UserStatusID])
}

func (s *UserService) translateUsers(
	userModels []models.User,
	userStatusesMap map[int64]string,
) []schemas.ResponseUser {

	responseUsers := make([]schemas.ResponseUser, 0, len(userModels))

	for _, u := range userModels {
		responseUsers = append(
			responseUsers,
			*s.transformUserModelIntoResponseUser(&u, userStatusesMap[u.UserStatusID]),
		)
	}

	return responseUsers
}
