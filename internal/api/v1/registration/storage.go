package registration

import (
	"context"

	"github.com/Alexander1000/service-auth/internal/model"
)

type storageRepository interface {
	Registration(ctx context.Context, userID int64, pass string, credentials []model.Credential) error
}
