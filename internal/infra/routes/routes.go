package routes

import (
	author_routes "github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author/routes"
	healthcheck_routes "github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/healthcheck/routes"
	metric_router "github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/observability/routes"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	serverMux      *chi.Mux
	registerRouter RegisterRoutes
	authorRouters  author_routes.Routes
}

func NewRoutes(
	serverMux *chi.Mux,
	registerRouter RegisterRoutes,
	authorRouters author_routes.Routes,
) *Routes {
	return &Routes{
		serverMux:      serverMux,
		registerRouter: registerRouter,
		authorRouters:  authorRouters,
	}
}
func (r *Routes) RegisterRoutes() {
	r.registerRouter.Register(r.serverMux, r.authorRouters.GroupRoutes())
	r.registerRouter.Register(r.serverMux, healthcheck_routes.NewHealthCheckRoutes().HealthCheckRoutes())
	r.registerRouter.Register(r.serverMux, metric_router.NewMetricRouter().MetricRoutes())
}
