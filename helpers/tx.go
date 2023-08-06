package helpers

import (
	"context"
	"database/sql"
)

func RollbackOrCommit(ctx context.Context, tx *sql.Tx) {
	defer func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				return
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				return
			}
		}
	}()
}
