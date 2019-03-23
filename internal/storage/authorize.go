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
			where token = '%s'`,
			token,
		),
	)

	err = row.Scan(&statusID, &expireAt)
	if err != nil {
		return err
	}

	if !statusID.Valid || !expireAt.Valid {
		return errors.New("invalid parse fields")
	}

	return nil
}
