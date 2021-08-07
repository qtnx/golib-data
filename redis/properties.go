package redis

import (
	_ "gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib/config"
)

type Properties struct {
	Host     string `default:"localhost"`
	Port     int    `default:"6379"`
	Database int    `default:"0"`
	Password string
}

func NewProperties(loader config.Loader) (*Properties, error) {
	props := Properties{}
	if err := loader.Bind(&props); err != nil {
		return nil, err
	}
	return &props, nil
}

func (p Properties) Prefix() string {
	return "application.redis"
}
