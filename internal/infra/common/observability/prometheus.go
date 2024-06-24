package observability

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

const (
	METER_NAME    = "foodtosave-case"
	METER_VERSION = "1.0.0"
)

type PrometheusAdapter struct {
	counterRequestHttpStatusCode      api.Int64Counter
	counterRedisError                 api.Int64Counter
	histogramRequestDuration          api.Float64Histogram
	histogramInstructionTableDuration api.Float64Histogram
	histogramOperationDuration        api.Float64Histogram
}

func NewPrometheusAdapter() *PrometheusAdapter {
	exporter, _ := prometheus.New()
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter(METER_NAME, api.WithInstrumentationVersion(METER_VERSION))
	counterRequestHttpStatusCode, _ := meter.Int64Counter("proxy_requests_total",
		api.WithDescription("Total number of proxy requests"))
	counterRedisError, _ := meter.Int64Counter("redis_errors_total",
		api.WithDescription("Total number of redis errors"))
	histogramRequestDuration, _ := meter.Float64Histogram("request_duration_seconds",
		api.WithDescription("Request duration in seconds"),
		api.WithExplicitBucketBoundaries(5, 10, 15, 20, 30, 50, 100, 200, 300, 500, 1000, 2000, 5000, 10000, 20000, 30000))
	histogramInstructionTableDuration, _ := meter.Float64Histogram("instruction_table_duration_seconds",
		api.WithDescription("Instruction table duration in seconds"),
		api.WithExplicitBucketBoundaries(5, 10, 15, 20, 30, 50, 100, 200, 300, 500, 1000, 2000, 5000, 10000, 20000, 30000))
	histogramOperationDuration, _ := meter.Float64Histogram("operation_duration_seconds",
		api.WithDescription("Operation duration in seconds"),
		api.WithExplicitBucketBoundaries(5, 10, 15, 20, 30, 50, 100, 200, 300, 500, 1000, 2000, 5000, 10000, 20000, 30000))

	return &PrometheusAdapter{
		counterRequestHttpStatusCode:      counterRequestHttpStatusCode,
		counterRedisError:                 counterRedisError,
		histogramRequestDuration:          histogramRequestDuration,
		histogramInstructionTableDuration: histogramInstructionTableDuration,
		histogramOperationDuration:        histogramOperationDuration,
	}
}
func (p *PrometheusAdapter) CounterRequestHttpStatusCode(ctx context.Context, router string, statusCode int) {
	opt := api.WithAttributes(
		attribute.Key("router").String(router),
		attribute.Key("status_code").Int(statusCode),
	)
	p.counterRequestHttpStatusCode.Add(ctx, 1, opt)
}
func (p *PrometheusAdapter) CounterRedisError(ctx context.Context, method string, codeError string) {
	opt := api.WithAttributes(
		attribute.Key("method").String(method),
		attribute.Key("code_error").String(codeError),
	)
	p.counterRedisError.Add(ctx, 1, opt)
}
func (p *PrometheusAdapter) HistogramRequestDuration(ctx context.Context, router string, statusCode int, duration float64) {
	opt := api.WithAttributes(
		attribute.Key("router").String(router),
		attribute.Key("status_code").Int(statusCode),
	)
	p.histogramRequestDuration.Record(ctx, duration, opt)
}
func (p *PrometheusAdapter) HistogramInstructionTableDuration(ctx context.Context, database, table, method string, duration float64) {
	opt := api.WithAttributes(
		attribute.Key("database").String(database),
		attribute.Key("table").String(table),
		attribute.Key("method").String(method),
	)
	p.histogramInstructionTableDuration.Record(ctx, duration, opt)
}

func (p *PrometheusAdapter) HistogramOperationDuration(ctx context.Context, operation string, context string, duration float64) {
	opt := api.WithAttributes(
		attribute.Key("operation").String(operation),
		attribute.Key("context").String(context),
	)
	p.histogramOperationDuration.Record(ctx, duration, opt)
}
