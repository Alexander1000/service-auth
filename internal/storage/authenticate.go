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

func (r *Repository) Authenticate(ctx context.Context, cred model.Credential, pass string) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if err != nil {
		return err
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
		return err
	}

	if !authID.Valid || authID.Int64 <= 0 {
		tx.Rollback()
		return errors.New("invalid auth_id")
	}
	if !passHash.Valid || len(passHash.String) == 0 {
		tx.Rollback()
		return errors.New("invalid pass_hash")
	}
	if !passSalt.Valid || len(passSalt.String) == 0 {
		tx.Rollback()
		return errors.New("invalid pass_salt")
	}

	u, err := uuid.Parse(passSalt.String)
	if err != nil {
		tx.Rollback()
		return err
	}
	uuidData, err := u.MarshalBinary()
	if err != nil {
		tx.Rollback()
		return err
	}

	hash := blake2b.Sum512(append([]byte(pass), uuidData...))
	if passHash.String != hex.EncodeToString(hash[0:]) {
		tx.Rollback()
		return errors.New("authenticate error")
	}

	var tokenID sql.NullInt64

	row = tx.QueryRowContext(
		ctx,
		fmt.Sprintf(`
			insert into auth_tokens(auth_id, token, status_id, created_at, expire_at)
			values (%d, '%s', %d, now(), now() + interval '1 day')
			returning token_id`,
			authID.Int64,
			uuid.New().String(), // todo refactoring random generate token
			0, // todo put in constants
		),
	)

	err = row.Scan(&tokenID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if !tokenID.Valid || tokenID.Int64 <= 0 {
		tx.Rollback()
		return errors.New("invalid token_id")
	}

	// todo: generate tokens and insert into auth_refresh_tokens

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
