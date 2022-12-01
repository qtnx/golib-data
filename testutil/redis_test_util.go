package golibdataTestUtil

import "github.com/go-redis/redis/v8"

var redisClient *redis.Client

func RedisClient() *redis.Client {
	return redisClient
}
