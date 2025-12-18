package mapper

import (
	"ticket-io/internal/user/domain"
	"ticket-io/internal/user/dto"
	"time"
)

func FormatUserToResponseUser(u *domain.User, statusName string) *dto.ResponseUser {

	return &dto.ResponseUser{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Birthdate: u.Birthdate.Format(time.RFC3339),
		Status:    statusName,
	}
}

func UserToResponseUser(u *domain.User, statusMap map[int64]string) *dto.ResponseUser {

	return FormatUserToResponseUser(u, statusMap[u.StatusID])
}

func UsersToResponseUsers(users []domain.User, statusMap map[int64]string) []dto.ResponseUser {

	formattedUsers := make([]dto.ResponseUser, 0, len(users))

	for _, u := range users {
		formattedUsers = append(formattedUsers, *FormatUserToResponseUser(&u, statusMap[u.StatusID]))
	}

	return formattedUsers
}
