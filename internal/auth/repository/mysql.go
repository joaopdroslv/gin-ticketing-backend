package repository

import (
	"context"
	"database/sql"
	"ticket-io/internal/auth/domain"
)

type mysqlAuthUserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *mysqlAuthUserRepository {

	return &mysqlAuthUserRepository{db: db}
}

func (r *mysqlAuthUserRepository) GetByEmail(ctx context.Context, email string) (*domain.AuthUser, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, email, password_hash FROM users WHERE email = ?
	`, email)

	var u domain.AuthUser

	if err := row.Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
	); err != nil {
		return nil, err
	}

	return &u, nil
}
