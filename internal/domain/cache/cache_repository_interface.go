package cache

type ICache interface {
	Set(key string, value any, ttl int)
	Get(key string) string
}
