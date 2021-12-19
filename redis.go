package golibdata

import (
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-data/redis"
	"go.uber.org/fx"
)

func RedisOpt() fx.Option {
	return fx.Options(
		golib.ProvideProps(redis.NewProperties),
		golib.ProvideHealthChecker(redis.NewHealthChecker),
		fx.Provide(redis.NewClient),
	)
}
