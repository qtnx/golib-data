package golibdata

import (
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-data/redis"
	"go.uber.org/fx"
)

func RedisOpt() fx.Option {
	return fx.Options(
		golib.ProvideProps(redis.NewProperties),
		golib.ProvideHealthChecker(redis.NewHealthChecker),
		fx.Provide(redis.NewClient),
	)
}
