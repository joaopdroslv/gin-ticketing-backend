package user

import "context"

type StatusProvider interface {
	GetStatusMap(ctx context.Context) (map[int64]string, error)
}
