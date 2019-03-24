package authorize

import (
	"context"
)

type storageRepository interface {
	Authorize(ctx context.Context, token string) (int, error)
}
