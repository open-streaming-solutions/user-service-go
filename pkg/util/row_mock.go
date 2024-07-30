package util

import (
	"github.com/jackc/pgx/v5"
	"github.com/open-streaming-solutions/user-service/internal/errors"
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
