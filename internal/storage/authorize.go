package storage

import (
	"context"
	"fmt"
	"database/sql"
	"errors"
)

func (r *Repository) Authorize(ctx context.Context, token string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	var statusID sql.NullInt64
	var expireAt sql.NullString

	row := conn.QueryRowContext(
		ctx,
		fmt.Sprintf(`
			select status_id, expire_at
			from auth_tokens
			where status_id in (0, 1) and token = '%s'`,
			token,
		),
	)

	err = row.Scan(&statusID, &expireAt)
	if err != nil {
		return err
	}

	if !statusID.Valid {
		return errors.New("invalid field status_id")
	}

	return nil
}
