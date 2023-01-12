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
	Host         string `validate:"required" default:"localhost"`
	Port         int    `validate:"required" default:"6379"`
	Database     int    `default:"0"`
	Username     string
	Password     string
	EnableTLS    bool
	PoolSize     int           `default:"10"`
	MinIdleConns int           `default:"2"`
	MaxConnAge   time.Duration `default:"0"` // Zero is to not close aged connections
	IdleTimeout  time.Duration `default:"5m"`
}

func (p Properties) Prefix() string {
	return "app.redis"
}
