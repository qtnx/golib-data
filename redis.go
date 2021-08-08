package golibdata

import (
	"fmt"
	red "github.com/go-redis/redis/v8"
	"gitlab.id.vin/vincart/golib-data/redis"
	"gitlab.id.vin/vincart/golib/config"
	"gitlab.id.vin/vincart/golib/web/actuator"
	"go.uber.org/fx"
)

type RedisOut struct {
	fx.Out
	Client        *red.Client
	HealthChecker actuator.HealthChecker `group:"actuator_health_checker"`
}

func NewRedisAutoConfig(loader config.Loader) (RedisOut, error) {
	out := RedisOut{}
	properties, err := redis.NewProperties(loader)
	if err != nil {
		return out, err
	}
	client, err := redis.NewClient(properties)
	if err != nil {
		return out, fmt.Errorf("cannot init redis with error [%s]", err.Error())
	}
	out.Client = client
	out.HealthChecker = redis.NewHealthChecker(client)
	return out, nil
}
