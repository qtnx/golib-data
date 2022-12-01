package golibdataTestUtil

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

func DatabaseTestUtilOpt() fx.Option {
	return fx.Provide(EnableDatabaseTestUtil)
}

func RedisTestUtilOpt() fx.Option {
	return fx.Invoke(func(rc *redis.Client) {
		redisClient = rc
	})
}

func TruncateTablesOpt(tables ...string) fx.Option {
	return fx.Invoke(func() {
		TruncateTables(tables...)
	})
}
