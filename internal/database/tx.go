package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/open-streaming-solutions/user-service/internal/errors"
	"github.com/open-streaming-solutions/user-service/pkg/util"
)

type Tx struct {
	tx pgx.Tx
}

func (tx *Tx) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	cmd, err := tx.tx.Exec(ctx, sql, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.ConvertPgError(err)
	}
	return cmd, nil
}

func (tx *Tx) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	rows, err := tx.tx.Query(ctx, sql, args...)
	if err != nil {
		return rows, errors.ConvertPgError(err)
	}
	return rows, nil
}

func (tx *Tx) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return util.MockRow{Row: tx.tx.QueryRow(ctx, sql, args...)}
}

func (tx *Tx) Begin(ctx context.Context) (pgx.Tx, error) {
	return tx.tx.Begin(ctx)
}

func (tx *Tx) Commit(ctx context.Context) error {
	return tx.tx.Commit(ctx)
}

func (tx *Tx) Rollback(ctx context.Context) error {
	return tx.tx.Rollback(ctx)
}

func (tx *Tx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return tx.tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

func (tx *Tx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return tx.tx.SendBatch(ctx, b)
}

func (tx *Tx) LargeObjects() pgx.LargeObjects {
	return tx.tx.LargeObjects()
}

func (tx *Tx) Prepare(ctx context.Context, name string, sql string) (*pgconn.StatementDescription, error) {
	return tx.tx.Prepare(ctx, name, sql)
}

func (tx *Tx) Conn() *pgx.Conn {
	return tx.tx.Conn()
}
