package redis

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func NewClient(props *Properties) (*redis.Client, error) {
	if props == nil {
		return nil, errors.New("redis config is required")
	}
	if len(props.Host) == 0 {
		return nil, errors.New("redis host is required")
	}
	if props.Port <= 0 {
		return nil, errors.New("redis port is required")
	}
	config := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", props.Host, props.Port),
		Username: props.Username,
		Password: props.Password,
		DB:       props.Database,
	}
	if props.EnableTLS {
		config.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}
	rdb := redis.NewClient(config)
	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return nil, errors.WithMessage(err, "error when open redis connection")
	}
	return rdb, nil
}
