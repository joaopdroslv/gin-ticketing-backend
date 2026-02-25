package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"

	shareddoamin "go-gin-ticketing-backend/internal/shared/domain"
	"go-gin-ticketing-backend/internal/shared/errs"
	"go-gin-ticketing-backend/internal/user/dto"
	"go-gin-ticketing-backend/internal/user/models"
)

type UserRepositoryMysql struct {
	db *sql.DB
}

func NewUserRepositoryMysql(db *sql.DB) *UserRepositoryMysql {

	return &UserRepositoryMysql{db: db}
}

func (r *UserRepositoryMysql) GetAllUsers(
	ctx context.Context,
	pagination *shareddoamin.Pagination,
) ([]models.User, *int64, error) {

	rows, err := r.db.QueryContext(ctx, `
		SELECT
			users.id,
			users.user_credential_id,
			users.user_status_id,
			users.name,
			users.birthdate,
			user_credentials.email,
			users.created_at,
			users.updated_at,
			COUNT(*) OVER() AS total_count
		FROM main.users
		JOIN main.user_credentials ON user_credentials.id = users.user_credential_id
		ORDER BY users.id DESC
		LIMIT ?
		OFFSET ?
	`, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	var total int64

	for rows.Next() {
		var user models.User
		var totalCount int64

		if err := rows.Scan(
			&user.ID,
			&user.UserCredentialID,
			&user.UserStatusID,
			&user.Name,
			&user.Birthdate,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
			&totalCount,
		); err != nil {
			return nil, nil, err
		}

		if total == 0 {
			total = totalCount
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return users, &total, nil
}

func (r *UserRepositoryMysql) GetUserByID(ctx context.Context, id int64) (*models.User, error) {

	row := r.db.QueryRowContext(ctx, `
		SELECT
			users.id,
			users.user_credential_id,
			users.user_status_id,
			users.name,
			users.birthdate,
			user_credentials.email,
			users.created_at,
			users.updated_at
		FROM main.users
		JOIN main.user_credentials ON user_credentials.id = users.user_credential_id
		WHERE users.id = ?
	`, id)

	var user models.User

	if err := row.Scan(
		&user.ID,
		&user.UserCredentialID,
		&user.UserStatusID,
		&user.Name,
		&user.Birthdate,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrResourceNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryMysql) CreateUser(
	ctx context.Context,
	data *dto.CreateUserData,
) (*int64, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(
		ctx,
		`INSERT INTO main.user_credentials (email) VALUES (?)`,
		data.Email,
	)
	if err != nil {
		return nil, err
	}

	userCredentialID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	result, err = tx.ExecContext(
		ctx,
		`INSERT INTO main.users (
			user_credential_id,
			user_status_id,
			name,
			birthdate
		) VALUES (?, ?, ?, ?)`,
		userCredentialID,
		data.UserStatusID,
		data.Name,
		data.Birthdate,
	)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *UserRepositoryMysql) UpdateUserByID(
	ctx context.Context,
	id int64,
	data *dto.UpdateUserData,
) (*models.User, error) {

	query, args, err := r.formatUpdateUserQuery(id, data)
	if err != nil {
		return nil, err
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errs.ErrZeroRowsAffected
	}

	return r.GetUserByID(ctx, id)
}

func (r UserRepositoryMysql) formatUpdateUserQuery(
	id int64,
	data *dto.UpdateUserData,
) (string, []any, error) {

	userFields := []string{}
	userCredentialFields := []string{}
	args := []any{}

	if data.Name != nil {
		userFields = append(userFields, "users.name = ?")
		args = append(args, data.Name)
	}
	if data.Birthdate != nil {
		userFields = append(userFields, "users.birthdate = ?")
		args = append(args, data.Birthdate)
	}
	if data.Email != nil {
		userCredentialFields = append(userCredentialFields, "user_credentials.email = ?")
		args = append(args, data.Email)
	}

	if len(userFields) == 0 && len(userCredentialFields) == 0 {
		return "", nil, fmt.Errorf("update user: %w", errs.ErrNothingToUpdate)
	}

	setItems := []string{}
	setItems = append(setItems, userFields...)
	setItems = append(setItems, userCredentialFields...)

	query := `UPDATE main.users`

	if len(userCredentialFields) > 0 {
		query += ` JOIN main.user_credentials ON user_credentials.id = users.user_credential_id`
	}

	query += fmt.Sprintf(" SET %s WHERE users.id = ?", strings.Join(setItems, ", "))
	args = append(args, id)

	log.Println(query)

	return query, args, nil
}

func (r *UserRepositoryMysql) DeleteUserByID(ctx context.Context, id int64) (bool, error) {

	result, err := r.db.ExecContext(ctx, `DELETE FROM main.users WHERE users.id = ?`, id)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, errs.ErrZeroRowsAffected
	}

	return true, nil
}
