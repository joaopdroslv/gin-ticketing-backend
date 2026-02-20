package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"go-gin-ticketing-backend/internal/user/models"
)

type UserStatusRepositoryMysql struct {
	db *sql.DB
}

func NewUserStatusRepositoryMysql(db *sql.DB) *UserStatusRepositoryMysql {

	return &UserStatusRepositoryMysql{db: db}
}

func (r *UserStatusRepositoryMysql) ListUserStatuses(ctx context.Context) ([]models.UserStatus, error) {

	rows, err := r.db.QueryContext(ctx, `
		SELECT
			user_statuses.id,
			user_statuses.name,
			user_statuses.description,
			user_statuses.created_at,
			user_statuses.updated_at
		FROM main.user_statuses
		ORDER BY user_statuses.id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userStatuses := make([]models.UserStatus, 0)

	for rows.Next() {
		var userStatus models.UserStatus

		if err := rows.Scan(
			&userStatus.ID,
			&userStatus.Name,
			&userStatus.Description,
			&userStatus.CreatedAt,
			&userStatus.UpdatedAt,
		); err != nil {
			return nil, err
		}

		userStatuses = append(userStatuses, userStatus)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userStatuses, nil
}
