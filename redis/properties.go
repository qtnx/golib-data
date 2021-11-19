package redis

import (
	_ "gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib/config"
)

func NewProperties(loader config.Loader) (*Properties, error) {
	props := Properties{}
	err := loader.Bind(&props)
	return &props, err
}

type Properties struct {
	Host     string `default:"localhost"`
	Port     int    `default:"6379"`
	Database int    `default:"0"`
	Password string
}

func (p Properties) Prefix() string {
	return "app.redis"
}
