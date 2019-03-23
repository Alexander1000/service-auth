package storage

import (
	"context"
	"github.com/Alexander1000/service-auth/internal/model"
	"fmt"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/blake2b"
	"encoding/hex"
)

func (r *Repository) Authenticate(ctx context.Context, cred model.Credential, pass string) (*model.Token, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(
		ctx,
		fmt.Sprintf(
			`select uac.auth_id, up.pass_hash, up.pass_salt
			from users_auth_credentials uac
			left join users_pass up on up.user_id = uac.user_id
			where uac.credential_id = %d and uac.type_id = %d`,
			cred.ID,
			r.credTypeMap[cred.Type],
		),
	)

	var authID sql.NullInt64
	var passHash, passSalt sql.NullString

	err = row.Scan(&authID, &passHash, &passSalt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !authID.Valid || authID.Int64 <= 0 {
		tx.Rollback()
		return nil, errors.New("invalid auth_id")
	}
	if !passHash.Valid || len(passHash.String) == 0 {
		tx.Rollback()
		return nil, errors.New("invalid pass_hash")
	}
	if !passSalt.Valid || len(passSalt.String) == 0 {
		tx.Rollback()
		return nil, errors.New("invalid pass_salt")
	}

	u, err := uuid.Parse(passSalt.String)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	uuidData, err := u.MarshalBinary()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	hash := blake2b.Sum512(append([]byte(pass), uuidData...))
	if passHash.String != hex.EncodeToString(hash[0:]) {
		tx.Rollback()
		return nil, errors.New("authenticate error")
	}

	var tokenID sql.NullInt64

	// todo refactoring random generate token
	authToken := uuid.New().String()

	row = tx.QueryRowContext(
		ctx,
		fmt.Sprintf(`
			insert into auth_tokens(auth_id, token, status_id, created_at, expire_at)
			values (%d, '%s', %d, now(), now() + interval '2 day')
			returning token_id`,
			authID.Int64,
			authToken,
			0, // todo put in constants
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

	// todo refactoring random generate token
	authRefreshToken := uuid.New().String()

	_, err = tx.ExecContext(
		ctx,
		fmt.Sprintf(`
			insert into auth_refresh_tokens(token_id, status_id, created_at, token, expire_at)
			values(%d, %d, now(), '%s', now() + interval '1 month')`,
			tokenID.Int64,
			0,
			authRefreshToken,
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

	return &model.Token{AccessToken: authToken, RefreshToken: authRefreshToken}, nil
}
