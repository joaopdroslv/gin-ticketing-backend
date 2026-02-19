package userstatus

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"go-gin-ticketing-backend/internal/user/models"
)

type mysqlUserStatusRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *mysqlUserStatusRepository {

	return &mysqlUserStatusRepository{db: db}
}

func (r *mysqlUserStatusRepository) ListUserStatuses(ctx context.Context) ([]models.UserStatus, error) {

	rows, err := r.db.QueryContext(ctx, `SELECT * FROM main.user_statuses ORDER BY id DESC`)
	if err != nil {
		return nil, fmt.Errorf("list user statuses query: %w", err)
	}
	defer rows.Close()

	userStatuses := make([]models.UserStatus, 0)

	for rows.Next() {
		var s models.UserStatus

		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Description,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("list user statuses scan: %w", err)
		}
		userStatuses = append(userStatuses, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("list user statuses rows error: %w", err)
	}

	return userStatuses, nil
}
