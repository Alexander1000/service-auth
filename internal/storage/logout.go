package storage

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) Logout(ctx context.Context, token string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if err != nil {
		return err
	}

	var dbTokenID, dbRefreshTokenID sql.NullInt64

	err = tx.QueryRowContext(
		ctx,
		fmt.Sprintf(`
			select
				at.token_id,
				at.status_id,
				at.expire_at,
				art.refresh_token_id,
				art.status_id,
				art.expire_at
			from auth_tokens at
			left join auth_refresh_tokens art on art.token_id = at.token_id
			where at.token = '%s'`,
			token,
		),
	).Scan(&dbTokenID, &dbRefreshTokenID)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
