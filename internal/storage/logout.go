package storage

import (
	"context"
)

func (r *Repository) Logout(ctx context.Context, token string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}
