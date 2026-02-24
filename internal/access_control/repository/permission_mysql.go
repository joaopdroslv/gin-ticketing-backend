package repository

import (
	"context"
	"database/sql"
	"go-gin-ticketing-backend/internal/access_control/models"
)

type PermissionRepositoryMysql struct {
	db *sql.DB
}

func NewPermissionRepositoryMysql(db *sql.DB) *PermissionRepositoryMysql {

	return &PermissionRepositoryMysql{db: db}
}

func (r *PermissionRepositoryMysql) GetPermissionsByUserID(ctx context.Context, id int64) ([]models.Permission, error) {

	rows, err := r.db.QueryContext(ctx, `
		SELECT
			permissions.id,
			permissions.name,
			permissions.created_at,
			permissions.updated_at
		FROM main.permissions
		JOIN main.role_permissions ON role_permissions.permission_id = permissions.id
		JOIN main.user_roles ON user_roles.role_id = role_permissions.role_id
		JOIN main.users ON users.id = user_roles.user_id
		WHERE users.id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []models.Permission

	for rows.Next() {
		var permission models.Permission

		if err := rows.Scan(
			&permission.ID,
			&permission.Name,
			&permission.CreatedAt,
			&permission.UpdatedAt,
		); err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	return permissions, nil
}

func (r *PermissionRepositoryMysql) UserHasPermission(
	ctx context.Context,
	id int64,
	permission string,
) (bool, error) {

	var exists int64

	err := r.db.QueryRowContext(ctx, `
		SELECT 1
		FROM main.users
		JOIN main.user_roles ON user_roles.user_id = users.id
		JOIN main.role_permissions ON role_permissions.role_id = user_roles.role_id
		JOIN main.permissions ON permissions.id = role_permissions.permission_id
		WHERE TRUE
			AND users.id = ?
			AND permissions = '?'
		LIMIT 1
	`, id, permission).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
