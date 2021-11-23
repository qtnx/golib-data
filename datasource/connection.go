package datasource

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.id.vin/vincart/golib-data/datasource/dialector"
	"gorm.io/gorm"
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
	connection, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, errors.WithMessage(err, "error when open connection")
	}
	sqlDb, err := connection.DB()
	if err != nil {
		return nil, errors.WithMessage(err, "datasource connection is not available")
	}
	// Config connection pool
	sqlDb.SetMaxIdleConns(p.MaxIdleConns)
	sqlDb.SetMaxOpenConns(p.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(p.ConnMaxLifetime)
	return connection, nil
}
