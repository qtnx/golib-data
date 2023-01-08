package redis

import (
	"github.com/go-redis/redis/v8"
	"gitlab.com/golibs-starter/golib/actuator"
)

type Informer struct {
	client *redis.Client
}

func NewInformer(client *redis.Client) actuator.Informer {
	return &Informer{client: client}
}

func (d Informer) Key() string {
	return "redis"
}

func (d Informer) Value() interface{} {
	return map[string]interface{}{
		"pool": map[string]interface{}{
			"total_conns": d.client.PoolStats().TotalConns,
			"idle_conns":  d.client.PoolStats().IdleConns,
			"stale_conns": d.client.PoolStats().StaleConns,
			"hits":        d.client.PoolStats().Hits,
			"misses":      d.client.PoolStats().Misses,
			"timeouts":    d.client.PoolStats().Timeouts,
		},
	}
}
