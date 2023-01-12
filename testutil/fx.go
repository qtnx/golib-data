package golibdataTestUtil

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// EnableDatabaseTestUtil
// Deprecated use DatabaseTestUtilOpt instead
func EnableDatabaseTestUtil() fx.Option {
	return DatabaseTestUtilOpt()
}

func DatabaseTestUtilOpt() fx.Option {
	return fx.Invoke(func(db *gorm.DB) {
		orm = db
	})
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
