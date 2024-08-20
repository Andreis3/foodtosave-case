package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
)

type Cache struct {
	client  *redis.Client
	metrics observability.IMetricAdapter
	log     logger.ILogger
}

func NewCache(client *redis.Client, metrics observability.IMetricAdapter, log logger.ILogger) *Cache {
	return &Cache{client: client, metrics: metrics, log: log}
}

func (c *Cache) Set(key string, value any, ttl int) {
	start := time.Now()
	ctx := context.Background()
	timeDuration := time.Duration(ttl) * time.Second
	valueMarshal, errM := json.Marshal(value)
	if errM != nil {
		c.metrics.CounterRedisError(ctx, "marshal", "RC-500")
		c.log.ErrorJson("Cache Marshal error", "error", errM.Error())
	}
	err := c.client.Set(ctx, key, string(valueMarshal), timeDuration).Err()
	if err != nil {
		c.metrics.CounterRedisError(ctx, "set", "RC-500")
		c.log.ErrorJson("Cache Set error", "error", err.Error(), "code", "RC-500", "origin", "Cache.Set")
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	c.metrics.HistogramInstructionTableDuration(ctx, "redis", "author-cache", "set", duration)
}
func (c *Cache) Get(key string) string {
	start := time.Now()
	ctx := context.Background()
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		c.metrics.CounterRedisError(ctx, "get", "RC-404")
		c.log.ErrorJson("Cache Get NotificationErrors", "error", "Key not found", "code", "RC-404", "origin", "Cache.Get")
		return ""
	} else if err != nil {
		c.metrics.CounterRedisError(ctx, "get", "RC-500")
		c.log.ErrorJson("Cache Get error", "error", err.Error(), "code", "RC-500", "origin", "Cache.Get")
		return ""
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	c.metrics.HistogramInstructionTableDuration(ctx, "redis", "author-cache", "get", duration)
	return val
}
