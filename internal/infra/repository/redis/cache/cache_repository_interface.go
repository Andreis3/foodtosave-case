package cache

import "github.com/andreis3/foodtosave-case/internal/util"

type ICache interface {
	Set(key string, value any, ttl int) *util.ValidationError
	Get(key string) (any, *util.ValidationError)
}
