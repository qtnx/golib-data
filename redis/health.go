package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gitlab.id.vin/vincart/golib/actuator"
	"gitlab.id.vin/vincart/golib/log"
)

type HealthChecker struct {
	client *redis.Client
}

func NewHealthChecker(client *redis.Client) actuator.HealthChecker {
	return &HealthChecker{client: client}
}

func (h HealthChecker) Component() string {
	return "redis"
}

func (h HealthChecker) Check() actuator.StatusDetails {
	statusDetails := actuator.StatusDetails{
		Status: actuator.StatusUp,
	}
	_, err := h.client.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Redis health check failed, err [%s]", err.Error())
		statusDetails.Status = actuator.StatusDown
		statusDetails.Reason = "Redis health check failed"
	}
	return statusDetails
}
