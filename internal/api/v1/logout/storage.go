package logout

import (
	"context"
)

type storageRepository interface {
	Logout(ctx context.Context, token string) error
}
