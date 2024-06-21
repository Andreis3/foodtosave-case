package cache

import (
	"context"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	client  *redis.Client
	metrics observability.IMetricAdapter
}

func NewCache(client *redis.Client, metrics observability.IMetricAdapter) *Cache {
	return &Cache{client: client}
}

func (c *Cache) Set(key string, value any, ttl int) *util.ValidationError {
	start := time.Now()
	ctx := context.Background()
	timeDuration := time.Duration(ttl) * time.Second
	err := c.client.Set(ctx, key, value, timeDuration).Err()
	if err != nil {
		return &util.ValidationError{
			Code:        "RC-500",
			Origin:      "Cache.Set",
			Status:      500,
			ClientError: []string{"Internal Server Error"},
			LogError:    []string{err.Error()},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	c.metrics.HistogramInstructionTableDuration(ctx, "redis", "author", "set", duration)
	return nil
}
func (c *Cache) Get(key string) (any, *util.ValidationError) {
	start := time.Now()
	ctx := context.Background()
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, &util.ValidationError{
			Code:        "RC-404",
			Origin:      "Cache.Get",
			Status:      404,
			ClientError: []string{"Not Found"},
			LogError:    []string{"Key not found"},
		}
	} else if err != nil {
		return nil, &util.ValidationError{
			Code:        "RC-501",
			Origin:      "Cache.Get",
			Status:      500,
			ClientError: []string{"Internal Server Error"},
			LogError:    []string{err.Error()},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	c.metrics.HistogramInstructionTableDuration(ctx, "redis", "author", "get", duration)
	return val, nil
}
