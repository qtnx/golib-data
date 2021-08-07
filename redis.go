package golibdata

import (
	"fmt"
	red "github.com/go-redis/redis/v8"
	"gitlab.id.vin/vincart/golib-data/redis"
	"gitlab.id.vin/vincart/golib/config"
)

func NewRedisAutoConfig(loader config.Loader) *red.Client {
	properties := redis.NewProperties(loader)
	client, err := redis.NewClient(properties)
	if err != nil {
		panic(fmt.Sprintf("Cannot init redis with error [%s]", err.Error()))
	}
	return client
}
