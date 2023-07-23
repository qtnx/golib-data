package datasource

import (
	"context"
	"database/sql"
	"gitlab.com/golibs-starter/golib/actuator"
	"gitlab.com/golibs-starter/golib/log"
)

type HealthChecker struct {
	connection *sql.DB
}

func NewHealthChecker(connection *sql.DB) actuator.HealthChecker {
	return &HealthChecker{connection: connection}
}

func (h HealthChecker) Component() string {
	return "datasource"
}

func (h HealthChecker) Check(ctx context.Context) actuator.StatusDetails {
	statusDetails := actuator.StatusDetails{
		Status: actuator.StatusUp,
	}
	if err := h.connection.PingContext(ctx); err != nil {
		log.WithCtx(ctx).WithError(err).Error("Datasource health check failed")
		statusDetails.Status = actuator.StatusDown
		statusDetails.Reason = "Datasource health check failed"
	}
	return statusDetails
}
