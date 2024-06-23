package metricsrouter

import (
	"github.com/andreis3/foodtosave-case/internal/presentation/http/helpers"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsRoutes struct {
}

func NewMetricRouter() *MetricsRoutes {
	return &MetricsRoutes{}
}
func (r *MetricsRoutes) MetricRoutes() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/metrics",
			Controller:  promhttp.Handler(),
			Description: "Metrics Prometheus",
			Type:        helpers.HANDLER,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
