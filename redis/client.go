package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewClient(props *Properties) (*redis.Client, error) {
	if props == nil {
		return nil, errors.New("missing redis config")
	}
	if len(props.Host) == 0 {
		return nil, errors.New("missing redis host")
	}
	if props.Port <= 0 {
		return nil, errors.New("missing redis port")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", props.Host, props.Port),
		Password: props.Password,
		DB:       props.Database,
	})
	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
