package repository

import "go.uber.org/fx"

var Module = fx.Provide(
	fx.Annotate(
		New,
		fx.As(new(Querier)),
	),
)
