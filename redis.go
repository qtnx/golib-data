package golibdata

import (
	"fmt"
	red "github.com/go-redis/redis/v8"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-data/redis"
)

func NewRedisAutoConfig(app *golib.App) *red.Client {
	properties := &redis.Properties{}
	app.ConfigLoader.Bind(properties)

	client, err := redis.NewClient(properties)
	if err != nil {
		panic(fmt.Sprintf("Cannot init redis with error [%s]", err.Error()))
	}
	return client
}
