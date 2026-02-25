package service

import (
	"context"
	"go-gin-ticketing-backend/internal/user/models"
	"go-gin-ticketing-backend/internal/user/repository"
)

type UserStatusService struct {
	userStatusRepository repository.UserStatusRepository
}

func NewUserStatusService(userStatusRepository repository.UserStatusRepository) *UserStatusService {

	return &UserStatusService{userStatusRepository: userStatusRepository}
}

func (s *UserStatusService) GetAllUserStatuses(ctx context.Context) ([]models.UserStatus, error) {

	return s.userStatusRepository.GetAllUserStatuses(ctx)
}

func (s *UserStatusService) GetUserStatusesMap(ctx context.Context) (map[int64]string, error) {

	userStatuses, err := s.userStatusRepository.GetAllUserStatuses(ctx)
	if err != nil {
		return nil, err
	}

	mapping := make(map[int64]string, len(userStatuses))

	for _, st := range userStatuses {
		mapping[st.ID] = st.Name
	}

	return mapping, nil
}
