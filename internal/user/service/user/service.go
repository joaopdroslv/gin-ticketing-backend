package user

import (
	"context"
	shareddomain "go-gin-ticketing-backend/internal/shared/domain"
	sharedschemas "go-gin-ticketing-backend/internal/shared/schemas"
	"go-gin-ticketing-backend/internal/user/domain"
	userrepository "go-gin-ticketing-backend/internal/user/repository/user"
	"go-gin-ticketing-backend/internal/user/schemas"
	"go-gin-ticketing-backend/internal/user/utils"
	"time"
)

type UserStatusProvider interface {
	GetUserStatusesMap(ctx context.Context) (map[int64]string, error)
}

type UserService struct {
	userRepository  userrepository.UserRepository
	userStatusesMap map[int64]string
}

func New(
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
		Items: utils.DomainUsersToResponseUsers(users, s.userStatusesMap),
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

	return utils.DomainUserToResponseUser(user, s.userStatusesMap), nil
}

func (s *UserService) CreateUser(ctx context.Context, body schemas.CreateUserBody) (*schemas.ResponseUser, error) {

	birthdate, err := time.Parse("2006-01-02", body.Birthdate)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(body.UserStatusID, body.Email, body.Name, birthdate)
	if err != nil {
		return nil, err
	}

	user, err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return utils.DomainUserToResponseUser(user, s.userStatusesMap), nil
}

func (s *UserService) UpdateUserByID(ctx context.Context, id int64, data schemas.UpdateUserBody) (*schemas.ResponseUser, error) {

	user, err := s.userRepository.UpdateUserByID(ctx, id, data)
	if err != nil {
		return nil, err
	}

	return utils.DomainUserToResponseUser(user, s.userStatusesMap), nil
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
