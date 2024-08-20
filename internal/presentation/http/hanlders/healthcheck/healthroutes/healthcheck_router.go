package healthroutes

import (
	"net/http"

	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/healthcheck/healthhandler"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/helpers"
)

type HealthCheckRouter struct{}

func NewHealthCheckRoutes() *HealthCheckRouter {
	return &HealthCheckRouter{}
}
func (r *HealthCheckRouter) HealthCheckRoutes() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/healthcheck",
			Controller:  healthhandler.HealthCheck,
			Description: "Health Check",
			Type:        helpers.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
