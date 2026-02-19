package user

import "context"

type UserStatusProvider interface {
	GetUserStatusesMap(ctx context.Context) (map[int64]string, error)
}
