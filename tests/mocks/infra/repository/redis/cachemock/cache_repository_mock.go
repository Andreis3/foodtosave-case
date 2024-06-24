package cachemock

import (
	"github.com/stretchr/testify/mock"
)

type CacheMock struct {
	mock.Mock
}

func (c *CacheMock) Set(key string, value any, ttl int) {
	c.Called(key, value, ttl)
}
func (c *CacheMock) Get(key string) string {
	args := c.Called(key)
	return args.Get(0).(string)
}
