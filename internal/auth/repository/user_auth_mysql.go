package repository

import (
	"context"
	"database/sql"
	"go-gin-ticketing-backend/internal/auth/domain"
)

type UserAuthRepositoryMysql struct {
	db *sql.DB
}

func NewUserAuthRepositoryMysql(db *sql.DB) *UserAuthRepositoryMysql {

	return &UserAuthRepositoryMysql{db: db}
}

func (r *UserAuthRepositoryMysql) GetUserByEmail(ctx context.Context, email string) (*domain.UserAuth, error) {

	row := r.db.QueryRowContext(ctx, `
		SELECT
			users.id,
			users.user_status_id,
			users.email,
			users.password_hash
		FROM main.users
		WHERE users.email = ?
	`, email)

	var userAuth domain.UserAuth

	if err := row.Scan(
		&userAuth.ID,
		&userAuth.UserStatusID,
		&userAuth.Email,
		&userAuth.PasswordHash,
	); err != nil {
		return nil, err
	}

	return &userAuth, nil
}

func (r *UserAuthRepositoryMysql) RegisterUser(ctx context.Context, user *domain.UserAuth) (*domain.UserAuth, error) {

	res, err := r.db.ExecContext(ctx,
		`INSERT INTO main.users (
			users.user_status_id,
			users.name,
			users.birthdate,
			users.email,
			users.password_hash
		) VALUES (?, ?, ?, ?, ?)`,
		user.UserStatusID, user.Name, user.Birthdate, user.Email, user.PasswordHash,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}
