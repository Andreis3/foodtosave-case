package metric_router

import (
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsRoutes struct {
}

func NewMetricRouter() *MetricsRoutes {
	return &MetricsRoutes{}
}
func (r *MetricsRoutes) MetricRoutes() util.RouteType {
	return util.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/metrics",
			Controller:  promhttp.Handler(),
			Description: "Metrics Prometheus",
			Type:        util.HANDLER,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
