package repository

import (
	"context"
	"database/sql"
	"fmt"
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
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

// DEPRECATED: should create a record in users and user_credentials tables simultaneously
func (r *UserRepositoryMysql) CreateUser(
	ctx context.Context,
	data *dto.UserCreateData,
) (*int64, error) {

	result, err := r.db.ExecContext(ctx,
		`INSERT INTO users (
			users.user_status_id,
			users.name,
			users.birthdate
		) VALUES (?, ?, ?)`,
		data.UserStatusID,
		data.Name,
		data.Birthdate,
	)
	if err != nil {
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
	data *dto.UserUpdateData,
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
	data *dto.UserUpdateData,
) (string, []any, error) {

	fields := []string{}
	args := []any{}

	if data.Name != nil {
		fields = append(fields, "users.name = ?")
		args = append(args, data.Name)
	}

	if data.Email != nil {
		fields = append(fields, "users.email = ?")
		args = append(args, data.Email)
	}

	if data.Birthdate != nil {
		fields = append(fields, "users.birthdate = ?")
		args = append(args, data.Birthdate)
	}

	if len(fields) == 0 {
		return "", nil, fmt.Errorf("update user: %w", errs.ErrNothingToUpdate)
	}

	query := fmt.Sprintf(
		"UPDATE main.users SET %s WHERE users.id = ?",
		strings.Join(fields, ", "),
	)

	args = append(args, id)

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
