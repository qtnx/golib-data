package datasource

import (
	"database/sql"
	"gitlab.com/golibs-starter/golib/actuator"
)

type Informer struct {
	connection *sql.DB
}

func NewInformer(connection *sql.DB) actuator.Informer {
	return &Informer{connection: connection}
}

func (d Informer) Key() string {
	return "datasource"
}

func (d Informer) Value() interface{} {
	return map[string]interface{}{
		"pool": map[string]interface{}{
			"max_open_connections": d.connection.Stats().MaxOpenConnections,
			"open_connections":     d.connection.Stats().OpenConnections,
			"inuse":                d.connection.Stats().InUse,
			"idle":                 d.connection.Stats().Idle,
		},
		"counter": map[string]interface{}{
			"wait_count":               d.connection.Stats().WaitCount,
			"wait_duration":            d.connection.Stats().WaitDuration.String(),
			"max_idle_closed":          d.connection.Stats().MaxIdleClosed,
			"max_idle_time_closed":     d.connection.Stats().MaxIdleTimeClosed,
			"max_idle_lifetime_closed": d.connection.Stats().MaxLifetimeClosed,
		},
	}
}
