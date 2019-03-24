package storage

import (
	"context"
	"github.com/Alexander1000/service-auth/internal/model"
	"database/sql"
)

func (r *Repository) Refresh(ctx context.Context, token string) (*model.Token, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return nil, nil
}
