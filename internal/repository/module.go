package repository

import (
	"github.com/jackc/pgx/v5"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		New,
		fx.As(new(QuerierWithTrx)),
	),
)

type QuerierWithTrx interface {
	Querier
	WithTx(tx pgx.Tx) *Queries
}

var _ QuerierWithTrx = (*Queries)(nil)
