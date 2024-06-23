package cache

import (
	"context"
	"encoding/json"
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	client  *redis.Client
	metrics observability.IMetricAdapter
	log     logger.ILogger
}

func NewCache(client *redis.Client, metrics observability.IMetricAdapter, log logger.ILogger) *Cache {
	return &Cache{client: client, metrics: metrics, log: log}
}

func (c *Cache) SetNX(key string, value any, ttl int) {
	start := time.Now()
	ctx := context.Background()
	timeDuration := time.Duration(ttl) * time.Second
	valueMarshal, errM := json.Marshal(value)
	if errM != nil {
		c.log.ErrorJson("Cache Marshal NotificationErrors", "error", errM.Error())
	}
	err := c.client.SetNX(ctx, key, string(valueMarshal), timeDuration).Err()
	if err != nil {
		c.log.ErrorJson("Cache SetNX NotificationErrors", "error", err.Error(), "code", "RC-500", "origin", "Cache.SetNX")
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	c.metrics.HistogramInstructionTableDuration(ctx, "redis", "author", "set", duration)
}
func (c *Cache) Get(key string) (string, *util.ValidationError) {
	start := time.Now()
	ctx := context.Background()
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", &util.ValidationError{
			Code:        "RC-404",
			Origin:      "Cache.Get",
			Status:      404,
			ClientError: []string{"Not Found"},
			LogError:    []string{"Key not found"},
		}
	} else if err != nil {
		return "", &util.ValidationError{
			Code:        "RC-501",
			Origin:      "Cache.Get",
			Status:      500,
			ClientError: []string{"Internal Server NotificationErrors"},
			LogError:    []string{err.Error()},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	c.metrics.HistogramInstructionTableDuration(ctx, "redis", "author", "get", duration)
	return val, nil
}
