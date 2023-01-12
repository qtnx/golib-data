package datasource

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/golibs-starter/golib-data/datasource/dialector"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection(resolver *dialector.Resolver, p *Properties) (*gorm.DB, error) {
	if len(p.Driver) == 0 {
		return nil, errors.New("driver is required")
	}
	dialStrategy, err := resolver.Resolve(p.Driver)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("cannot resolve driver [%s]", p.Driver))
	}
	dial, err := dialStrategy.Open(&dialector.Config{
		Host:     p.Host,
		Port:     p.Port,
		Database: p.Database,
		Username: p.Username,
		Password: p.Password,
		Params:   p.Params,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "error when open dial")
	}
	gormLogLevel, err := getGormLogLevel(p.LogLevel)
	if err != nil {
		return nil, errors.WithMessage(err, "error when get log level")
	}
	connection, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		return nil, errors.WithMessage(err, "error when open connection")
	}
	sqlDb, err := connection.DB()
	if err != nil {
		return nil, errors.WithMessage(err, "datasource connection is not available")
	}
	// Config connection pool
	sqlDb.SetMaxOpenConns(p.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(p.ConnMaxLifetime)
	sqlDb.SetMaxIdleConns(p.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(p.ConnMaxIdleTime)
	return connection, nil
}

func getGormLogLevel(logLevel string) (logger.LogLevel, error) {
	switch logLevel {
	case "SILENT":
		return logger.Silent, nil
	case "ERROR":
		return logger.Error, nil
	case "WARN":
		return logger.Warn, nil
	case "INFO":
		return logger.Info, nil
	default:
		return 0, errors.New(fmt.Sprintf("log level [%s] is not valid", logLevel))
	}
}
