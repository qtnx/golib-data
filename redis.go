package golibdata

import (
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib-data/redis"
	"go.uber.org/fx"
)

func RedisOpt() fx.Option {
	return fx.Options(
		golib.ProvideProps(redis.NewProperties),
		golib.ProvideHealthChecker(redis.NewHealthChecker),
		golib.ProvideInformer(redis.NewInformer),
		fx.Provide(redis.NewClient),
	)
}
