package redis

import (
	"crypto/tls"
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
		return nil, err
	}
	return rdb, nil
}
