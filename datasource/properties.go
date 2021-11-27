package datasource

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
	Driver   string
	Host     string `default:"localhost"`
	Port     int    `default:"3306"`
	Database string
	Username string
	Password string
	Params   string

	// The maximum number of connections in the idle connection pool.
	MaxIdleConns int `default:"10"`

	// The maximum number of open connections to the database.
	MaxOpenConns int `default:"100"`

	// The maximum amount of time a connection may be reused.
	ConnMaxLifetime time.Duration `default:"30m"`
}

func (p Properties) Prefix() string {
	return "app.datasource"
}
