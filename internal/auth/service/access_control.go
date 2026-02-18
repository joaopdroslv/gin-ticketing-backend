package service

import "context"

type AccessControl interface {
	HasThisPermission(ctx context.Context, userID int64, permission string) (bool, error)
}
