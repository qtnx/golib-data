package golibdataTestUtil

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func EnableDatabaseTestUtil() fx.Option {
	return fx.Invoke(func(db *gorm.DB) {
		orm = db
	})
}

// DatabaseTestUtilOpt
// Deprecated: use EnableDatabaseTestUtil instead
func DatabaseTestUtilOpt() fx.Option {
	return EnableDatabaseTestUtil()
}

func EnableRedisTestUtilOpt() fx.Option {
	return fx.Invoke(func(rc *redis.Client) {
		redisClient = rc
	})
}

// RedisTestUtilOpt
// Deprecated: use EnableRedisTestUtilOpt instead
func RedisTestUtilOpt() fx.Option {
	return EnableRedisTestUtilOpt()
}

func TruncateTablesOpt(tables ...string) fx.Option {
	return fx.Invoke(func() {
		TruncateTables(tables...)
	})
}

func TruncateTablesHasForeignKeyOpt(tables ...string) fx.Option {
	return fx.Invoke(func() {
		TruncateTablesHasForeignKey(tables...)
	})
}
