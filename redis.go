package golibdata

import (
	"fmt"
	red "github.com/go-redis/redis/v8"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-data/redis"
	"gitlab.id.vin/vincart/golib/actuator"
	"go.uber.org/fx"
)

func RedisOpt() fx.Option {
	return fx.Options(
		golib.EnablePropsAutoload(new(redis.Properties)),
		fx.Provide(redis.NewProperties),
		fx.Provide(NewRedis),
	)
}

type RedisOut struct {
	fx.Out
	Client        *red.Client
	HealthChecker actuator.HealthChecker `group:"actuator_health_checker"`
}

func NewRedis(properties *redis.Properties) (RedisOut, error) {
	out := RedisOut{}
	client, err := redis.NewClient(properties)
	if err != nil {
		return out, fmt.Errorf("cannot init redis with error [%s]", err.Error())
	}
	out.Client = client
	out.HealthChecker = redis.NewHealthChecker(client)
	return out, nil
}
