package storage

import (
	"context"
	"github.com/Alexander1000/service-auth/internal/model"
)

func (r *Repository) Refresh(ctx context.Context, token string) (*model.Token, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	
	return nil, nil
}
