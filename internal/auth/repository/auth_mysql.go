package repository

import (
	"context"
	"database/sql"
	"go-gin-ticketing-backend/internal/auth/domain"
)

type AuthRepositoryMysql struct {
	db *sql.DB
}

func NewAuthRepositoryMysql(db *sql.DB) *AuthRepositoryMysql {

	return &AuthRepositoryMysql{db: db}
}

func (r *AuthRepositoryMysql) GetUserByEmail(
	ctx context.Context,
	email string,
) (*domain.UserCredential, error) {

	row := r.db.QueryRowContext(ctx, `
		SELECT
			user_credentials.email,
			user_credentials.password_hash,
			users.id,
			users.user_status_id,
			users.name,
			users.birthdate
		FROM main.user_credentials
		JOIN main.users ON users.user_credential_id = user_credentials.id
		WHERE user_credentials.email = ?
	`, email)

	var userCredential domain.UserCredential

	if err := row.Scan(
		&userCredential.Email,
		&userCredential.PasswordHash,
		&userCredential.UserInfo.ID,
		&userCredential.UserInfo.UserStatusID,
		&userCredential.UserInfo.Name,
		&userCredential.UserInfo.Birthdate,
	); err != nil {
		return nil, err
	}

	return &userCredential, nil
}

func (r *AuthRepositoryMysql) RegisterUser(
	ctx context.Context,
	userCredential *domain.UserCredential,
) (*domain.UserCredential, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(
		ctx,
		`INSERT INTO main.user_credentials (email, password_hash) VALUES (?, ?)`,
		userCredential.Email,
		userCredential.PasswordHash,
	)
	if err != nil {
		return nil, err
	}

	userCredentialID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	res, err = tx.ExecContext(
		ctx,
		`
		INSERT INTO main.users (
			user_credential_id,
			user_status_id,
			name,
			birthdate
		) VALUES (?, ?, ? ,?)
		`,
		userCredentialID,
		userCredential.UserInfo.UserStatusID,
		userCredential.UserInfo.Name,
		userCredential.UserInfo.Birthdate,
	)
	if err != nil {
		return nil, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	userCredential.ID = userCredentialID
	userCredential.UserInfo.ID = userID

	return userCredential, nil
}
