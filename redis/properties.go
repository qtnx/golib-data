package redis

import _ "gitlab.id.vin/vincart/golib"

type Properties struct {
	Host     string `default:"localhost"`
	Port     int    `default:"6379"`
	Database int    `default:"0"`
	Password string
}

func (p Properties) Prefix() string {
	return "application.redis"
}
