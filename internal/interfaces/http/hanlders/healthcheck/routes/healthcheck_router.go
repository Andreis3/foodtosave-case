package healthcheck_routes

import (
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/healthcheck/controller"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
)

type HealthCheckRouter struct{}

func NewHealthCheckRoutes() *HealthCheckRouter {
	return &HealthCheckRouter{}
}
func (r *HealthCheckRouter) HealthCheckRoutes() util.RouteType {
	return util.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/healthcheck",
			Controller:  healthcheck_controller.HealthCheck,
			Description: "Health Check",
			Type:        util.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
