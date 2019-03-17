package storage

import (
	"context"
	"github.com/Alexander1000/service-auth/internal/model"
	"fmt"
	"database/sql"
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

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
