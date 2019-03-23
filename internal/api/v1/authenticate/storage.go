package authenticate

import (
	"context"
	"github.com/Alexander1000/service-auth/internal/model"
)

type storageRepository interface {
	Authenticate(ctx context.Context, cred model.Credential, pass string) (*model.Token, error)
}

