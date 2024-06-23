package cache

import "github.com/andreis3/foodtosave-case/internal/util"

type ICache interface {
	SetNX(key string, value any, ttl int)
	Get(key string) (string, *util.ValidationError)
}
