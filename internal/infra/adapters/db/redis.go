package db

import (
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/configs"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(conf configs.Conf) *Redis {
	db, _ := strconv.Atoi(conf.RedisDB)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPassword,
		DB:       db,
	})

	return &Redis{
		client: client,
	}
}

func (r *Redis) InstanceDB() any {
	return r.client
}

func (r *Redis) Close() {
	r.client.Close()
}
