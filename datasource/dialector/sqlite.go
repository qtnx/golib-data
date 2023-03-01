package dialector

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
}

func NewSqlite() Strategy {
	return &Sqlite{}
}

func (m Sqlite) Driver() string {
	return "sqlite"
}

func (m Sqlite) Open(cf *Config) (gorm.Dialector, error) {
	dsn, err := m.buildDsn(cf)
	if err != nil {
		return nil, err
	}
	return sqlite.Open(dsn), nil
}

func (m Sqlite) buildDsn(cf *Config) (string, error) {
	if len(cf.Dsn) == 0 {
		return "", errors.New("DSN is required")
	}
	return cf.Dsn, nil
}
