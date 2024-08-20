package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/authorhandler/authorroutes"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/healthcheck/healthroutes"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/observability/metricsroutes"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/helpers"
)

type RegisterRoutes struct {
	serverMux     *chi.Mux
	authorRouters authorroutes.Routes
	logger        logger.ILogger
}

func NewRegisterRoutes(
	serverMux *chi.Mux,
	authorRouters authorroutes.Routes,
	logger logger.ILogger,
) *RegisterRoutes {
	return &RegisterRoutes{
		serverMux:     serverMux,
		authorRouters: authorRouters,
		logger:        logger,
	}
}
func (r *RegisterRoutes) RegisterRoutes() {
	r.register(r.serverMux, r.authorRouters.AuthorRoutes())
	r.register(r.serverMux, healthroutes.NewHealthCheckRoutes().HealthCheckRoutes())
	r.register(r.serverMux, metricsrouter.NewMetricRouter().MetricRoutes())
}

func (r *RegisterRoutes) register(serverMux *chi.Mux, router helpers.RouteType) {
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
