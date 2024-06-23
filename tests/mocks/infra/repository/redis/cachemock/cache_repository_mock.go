package cachemock

import (
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/stretchr/testify/mock"
)

type CacheMock struct {
	mock.Mock
}

func (c *CacheMock) SetNX(key string, value any, ttl int) {
	c.Called(key, value, ttl)
}
func (c *CacheMock) Get(key string) (string, *util.ValidationError) {
	args := c.Called(key)
	return args.Get(0).(string), args.Get(1).(*util.ValidationError)
}
