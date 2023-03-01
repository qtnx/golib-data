package dialector

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
}

func NewPostgres() Strategy {
	return &Postgres{}
}

func (p Postgres) Driver() string {
	return "postgres"
}

func (p Postgres) Open(cf *Config) (gorm.Dialector, error) {
	dsn, err := p.buildDsn(cf)
	if err != nil {
		return nil, err
	}
	return postgres.Open(dsn), nil
}

func (p Postgres) buildDsn(cf *Config) (string, error) {
	if len(cf.Dsn) > 0 {
		return cf.Dsn, nil
	}
	if len(cf.Host) == 0 {
		return "", errors.New("host is required")
	}
	if len(cf.Database) == 0 {
		return "", errors.New("database is required")
	}
	format := "host=%s user=%s password=%s dbname=%s port=%d %s"
	return fmt.Sprintf(format, cf.Host, cf.Username, cf.Password, cf.Database, cf.Port, cf.Params), nil
}
