package storage

import (
	"context"

	"github.com/Alexander1000/service-auth/internal/model"
	"database/sql"
	"fmt"
	"strings"
)

func (r *Repository) Registration(ctx context.Context, userID int64, pass string, credentials []model.Credential) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead, ReadOnly: false})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		fmt.Sprintf(`
			insert into users_pass(user_id, pass_hash, pass_salt, created_at)
			values (%d, %s, %s, now())`,
			userID,
		),
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	tuples := make([]string, 0, len(credentials))
	for _, cred := range credentials {
		tuples = append(
			tuples,
			fmt.Sprintf(
				"(%d, %d, %d, now())",
				userID,
				cred.ID,
				r.credTypeMap[cred.Type],
			),
		)
	}

	_, err = tx.ExecContext(
		ctx,
		fmt.Sprintf(`
			insert into users_auth_credentials(user_id, credential_id, type_id, created_at)
			values %s`,
			strings.Join(tuples, ", "),
		),
	)
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
