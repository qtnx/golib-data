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

func NewProperties(loader config.Loader) *Properties {
	props := Properties{}
	loader.Bind(&props)
	return &props
}

func (p Properties) Prefix() string {
	return "application.redis"
}
