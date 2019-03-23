package storage

import (
	"context"
	"fmt"
	"database/sql"
	"errors"
	"time"
)

func (r *Repository) Authorize(ctx context.Context, token string) (int, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return AuthInternalError, err
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
		if err == sql.ErrNoRows {
			return AuthNotFound, nil
		}
		return AuthInternalError, err
	}

	if !statusID.Valid || !expireAt.Valid {
		return AuthInternalError, errors.New("invalid parse fields")
	}

	if statusID.Int64 == AccessTokenStatusDisabled {
		return AuthDisabled, errors.New("token disabled")
	} else if statusID.Int64 == AccessTokenStatusRefreshed {
		return AuthRefreshed, errors.New("token refreshed")
	} else if statusID.Int64 != AccessTokenStatusActive {
		return AuthInternalError, errors.New("invalid status")
	}

	// active status
	expireTime, err := time.Parse(time.RFC3339, expireAt.String)
	if err != nil {
		return AuthInternalError, err
	}

	curTime := time.Now()

	if curTime.Unix() > expireTime.Unix() {
		return AuthExpired, errors.New("token expired")
	}

	return AuthOk, nil
}
