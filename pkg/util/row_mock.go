package util

import (
	"github.com/Open-Streaming-Solutions/user-service/internal/errors"
	"github.com/jackc/pgx/v5"
)

type MockRow struct {
	Row pgx.Row
}

func (r MockRow) Scan(dest ...any) error {
	if err := r.Row.Scan(dest...); err != nil {
		return errors.ConvertPgError(err)
	}
	return nil
}
