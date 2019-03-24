package storage

import (
	"context"
	"github.com/Alexander1000/service-auth/internal/model"
	"database/sql"
	"fmt"
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

	var refreshTokenID, tokenID, statusID sql.NullInt64
	var refreshExpireAt sql.NullString

	row := tx.QueryRowContext(
		ctx,
		fmt.Sprintf(`
			select refresh_token_id, token_id, status_id, expire_at
			from auth_refresh_tokens
			where token = '%s'`,
			token,
		),
	)

	err = row.Scan(&refreshTokenID, &tokenID, &statusID, &refreshExpireAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// todo check: expire, status

	_, err = tx.ExecContext(
		ctx,
		fmt.Sprintf(`
			update auth_refresh_tokens
			set status_id = %d, updated_at = now()
			where refresh_token_id = %d`,
			RefreshTokenStatusRefreshed,
			refreshTokenID.Int64,
		),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// todo check affected rows

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return nil, nil
}
