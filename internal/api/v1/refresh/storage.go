package refresh

import (
	"context"

	"github.com/Alexander1000/service-auth/internal/model"
)

type storageRepository interface {
	Refresh(ctx context.Context, token string) (*model.Token, error)
}
