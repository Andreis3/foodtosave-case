package routes

import (
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/authorroutes"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/healthcheck/healthroutes"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/observability/metricsroutes"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	serverMux      *chi.Mux
	registerRouter RegisterRoutes
	authorRouters  authorroutes.Routes
}

func NewRoutes(
	serverMux *chi.Mux,
	registerRouter RegisterRoutes,
	authorRouters authorroutes.Routes,
) *Routes {
	return &Routes{
		serverMux:      serverMux,
		registerRouter: registerRouter,
		authorRouters:  authorRouters,
	}
}
func (r *Routes) RegisterRoutes() {
	r.registerRouter.Register(r.serverMux, r.authorRouters.AuthorRoutes())
	r.registerRouter.Register(r.serverMux, healthroutes.NewHealthCheckRoutes().HealthCheckRoutes())
	r.registerRouter.Register(r.serverMux, metricsrouter.NewMetricRouter().MetricRoutes())
}
