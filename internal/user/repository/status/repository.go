package status

import (
	"context"
	"ticket-io/internal/user/domain"
)

type StatusRepository interface {
	ListStatuses(ctx context.Context) ([]domain.Status, error)
}
