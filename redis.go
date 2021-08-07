package golibdata

import (
	"fmt"
	red "github.com/go-redis/redis/v8"
	"gitlab.id.vin/vincart/golib-data/redis"
	"gitlab.id.vin/vincart/golib/config"
)

func NewRedisAutoConfig(loader config.Loader) (*red.Client, error) {
	properties, err := redis.NewProperties(loader)
	if err != nil {
		return nil, err
	}
	client, err := redis.NewClient(properties)
	if err != nil {
		return nil, fmt.Errorf("cannot init redis with error [%s]", err.Error())
	}
	return client, nil
}
