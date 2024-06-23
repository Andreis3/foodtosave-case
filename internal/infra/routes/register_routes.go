package routes

import (
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RegisterRoutes struct {
	logger logger.ILogger
}

func NewRegisterRoutes(logger logger.ILogger) *RegisterRoutes {
	return &RegisterRoutes{
		logger: logger,
	}
}
func (r *RegisterRoutes) Register(serverMux *chi.Mux, router helpers.RouteType) {
	message, info := "[RegisterRoutes] ", "MAPPED_ROUTER"
	for _, route := range router {
		switch route.Type {
		case helpers.HANDLER:
			switch len(route.Middlewares) > 0 {
			case true:
				methodAndPath := fmt.Sprintf("%s %s", route.Method, route.Path)
				r.logger.InfoText(message, info, fmt.Sprintf("%s - %s", methodAndPath, route.Description))
				serverMux.With(route.Middlewares...).Handle(methodAndPath, route.Controller.(http.Handler))
			default:
				methodAndPath := fmt.Sprintf("%s %s", route.Method, route.Path)
				r.logger.InfoText(message, info, fmt.Sprintf("%s - %s", methodAndPath, route.Description))
				serverMux.Handle(methodAndPath, route.Controller.(http.Handler))
			}

		case helpers.HANDLER_FUNC:
			switch len(route.Middlewares) > 0 {
			case true:
				methodAndPath := fmt.Sprintf("%s %s", route.Method, route.Path)
				r.logger.InfoText(message, info, fmt.Sprintf("%s - %s", methodAndPath, route.Description))
				serverMux.With(route.Middlewares...).HandleFunc(methodAndPath, route.Controller.(func(http.ResponseWriter, *http.Request)))
			default:
				methodAndPath := fmt.Sprintf("%s %s", route.Method, route.Path)
				r.logger.InfoText(message, info, fmt.Sprintf("%s %s - %s", route.Method, route.Path, route.Description))
				serverMux.HandleFunc(methodAndPath, route.Controller.(func(http.ResponseWriter, *http.Request)))
			}
		}
	}
}
