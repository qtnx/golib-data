package datasource

import (
	_ "github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib/config"
	"time"
)

func NewProperties(loader config.Loader) (*Properties, error) {
	props := Properties{}
	err := loader.Bind(&props)
	return &props, err
}

type Properties struct {
	Driver   string `validate:"required"`
	Dsn      string `validate:"required_without_all=Host Port Database Username"`
	Host     string `validate:"required_without=Dsn"`
	Port     int    `validate:"required_without=Dsn"`
	Database string `validate:"required_without=Dsn"`
	Username string `validate:"required_without=Dsn"`
	Password string
	Params   string

	// The maximum number of open connections to the database.
	MaxOpenConns int `default:"10"`

	// The maximum amount of time a connection may be reused.
	ConnMaxLifetime time.Duration `default:"30m"`

	// The maximum number of connections in the idle connection pool.
	MaxIdleConns int `default:"2"`

	// The maximum amount of time a connection may be idle.
	ConnMaxIdleTime time.Duration `default:"10m"`

	// The log level of gorm. There are four levels: "SILENT", "ERROR", "WARN", "INFO".
	LogLevel string `default:"SILENT"`
}

func (p Properties) Prefix() string {
	return "app.datasource"
}
