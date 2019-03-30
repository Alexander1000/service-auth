package storage

import (
	"context"
	"database/sql"
	"errors"
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
	var dbTokenStatus, dbRefreshTokenStatus sql.NullInt64
	var dbTokenExpire, dbRefreshTokenExpire sql.NullString

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
	).Scan(&dbTokenID, &dbTokenStatus, &dbTokenExpire, &dbRefreshTokenID, &dbRefreshTokenStatus, &dbRefreshTokenExpire)

	if err != nil {
		tx.Rollback()
		return err
	}

	if !dbTokenID.Valid {
		tx.Rollback()
		return errors.New("not found")
	}

	if dbTokenStatus.Int64 == AccessTokenStatusActive {
		_, err = tx.ExecContext(
			ctx,
			fmt.Sprintf(`
				update auth_tokens
				set status_id = %d, updated_at = now()
				where token_id = %d`,
				AccessTokenStatusDisabled,
				dbTokenID.Int64,
			),
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
