package dialector

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
}

func NewMysql() Strategy {
	return &Mysql{}
}

func (m Mysql) Driver() string {
	return "mysql"
}

func (m Mysql) Open(cf *Config) (gorm.Dialector, error) {
	dsn, err := m.buildDsn(cf)
	if err != nil {
		return nil, err
	}
	return mysql.Open(dsn), nil
}

func (m Mysql) buildDsn(cf *Config) (string, error) {
	if len(cf.Host) == 0 {
		return "", errors.New("host is required")
	}
	if len(cf.Database) == 0 {
		return "", errors.New("database is required")
	}
	params := cf.Params
	if len(params) > 0 {
		params = "?" + params
	}
	format := "%v:%v@tcp(%v:%v)/%v%v"
	return fmt.Sprintf(format, cf.Username, cf.Password, cf.Host, cf.Port, cf.Database, params), nil
}
