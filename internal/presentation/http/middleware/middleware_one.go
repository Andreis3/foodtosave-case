package middleware

import (
	"net/http"

	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
)

type MiddlewareOne struct {
	log logger.ILogger
}

func NewMiddlewareOne(log logger.ILogger) *MiddlewareOne {
	return &MiddlewareOne{
		log: log,
	}
}

func (m *MiddlewareOne) One(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.log.InfoText("middleware one")
		next.ServeHTTP(w, r)
	})
}
