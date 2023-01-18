package redis

import (
	_ "gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib/config"
	"time"
)

func NewProperties(loader config.Loader) (*Properties, error) {
	props := Properties{}
	err := loader.Bind(&props)
	return &props, err
}

type Properties struct {
	Host      string `validate:"required" default:"localhost"`
	Port      int    `validate:"required" default:"6379"`
	Database  int    `default:"0"`
	Username  string
	Password  string
	EnableTLS bool

	// Maximum number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	PoolSize int `default:"10"`

	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnAge time.Duration `default:"0"`

	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int `default:"2"`

	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration `default:"5m"`
}

func (p Properties) Prefix() string {
	return "app.redis"
}
