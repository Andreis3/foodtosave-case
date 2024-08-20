package middleware

import (
	"net/http"

	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
)

type MiddlewareTwo struct {
	log logger.ILogger
}

func NewMiddlewareTwo(log logger.ILogger) *MiddlewareTwo {
	return &MiddlewareTwo{
		log: log,
	}
}

func (m *MiddlewareTwo) Two(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.log.InfoText("middleware two")
		next.ServeHTTP(w, r)
	})
}
