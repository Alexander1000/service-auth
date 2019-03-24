package storage

import (
	"context"
	"github.com/Alexander1000/service-auth/internal/model"
	"database/sql"
	"fmt"
	"errors"
	"github.com/google/uuid"
	"time"
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

	var refreshTokenID, tokenID, refreshStatusID sql.NullInt64
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

	err = row.Scan(&refreshTokenID, &tokenID, &refreshStatusID, &refreshExpireAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !refreshTokenID.Valid || !tokenID.Valid || !refreshStatusID.Valid || !refreshExpireAt.Valid {
		tx.Rollback()
		return nil, errors.New("invalid parse")
	}

	if refreshStatusID.Int64 != RefreshTokenStatusActive {
		tx.Rollback()
		return nil, errors.New("not found")
	}

	tExpire, err := time.Parse(time.RFC3339, refreshExpireAt.String)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if tExpire.Unix() > time.Now().Unix() {
		tx.Rollback()
		return nil, errors.New("refresh token expired")
	}

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

	var authID, statusID sql.NullInt64
	var expireAt sql.NullString

	row = tx.QueryRowContext(
		ctx,
		fmt.Sprintf(`
			select auth_id, status_id, expire_at
			from auth_tokens
			where token_id = %d`,
			tokenID.Int64,
		),
	)

	err = row.Scan(&authID, &statusID, &expireAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !authID.Valid || !statusID.Valid || !expireAt.Valid {
		tx.Rollback()
		return nil, errors.New("invalid fields")
	}

	if statusID.Int64 != AccessTokenStatusActive {
		tx.Rollback()
		return nil, errors.New("invalid status auth token")
	}

	_, err = tx.ExecContext(
		ctx,
		fmt.Sprintf(`
			update auth_tokens
			set status_id = %d, updated_at = now()
			where token_id = %d`,
			AccessTokenStatusRefreshed,
			tokenID.Int64,
		),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	row = tx.QueryRowContext(
		ctx,
		fmt.Sprintf(`
			insert into auth_tokens(auth_id, token, status_id, created_at, expire_at)
			values (%d, '%s', %d, now(), now() + interval '2 day')
			returning token_id`,
			authID.Int64,
			uuid.New().String(),
			AccessTokenStatusActive,
		),
	)

	err = row.Scan(&tokenID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !tokenID.Valid || tokenID.Int64 <= 0 {
		tx.Rollback()
		return nil, errors.New("invalid token_id")
	}

	_, err = tx.ExecContext(
		ctx,
		fmt.Sprintf(`
			insert into auth_refresh_tokens(token_id, status_id, created_at, token, expire_at)
			values(%d, %d, now(), '%s', now() + interval '1 month')`,
			tokenID.Int64,
			RefreshTokenStatusActive,
			uuid.New().String(),
		),
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return nil, nil
}
