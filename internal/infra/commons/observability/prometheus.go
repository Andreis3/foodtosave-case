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
	counterConsumerContext           api.Int64Counter
	histogramConsumerContextDuration api.Float64Histogram
	histogramOperationTableDuration  api.Float64Histogram
}

func NewPrometheusAdapter() *PrometheusAdapter {
	exporter, _ := prometheus.New()
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter(METER_NAME, api.WithInstrumentationVersion(METER_VERSION))
	counterConsumerContext, _ := meter.Int64Counter("count_consumer_context",
		api.WithDescription("Count total number messages by context"))
	histogramConsumerContextDuration, _ := meter.Float64Histogram("consumer_duration_seconds",
		api.WithDescription("Consumer duration in seconds"),
		api.WithExplicitBucketBoundaries(5, 10, 15, 20, 30, 50, 100, 200, 300, 500, 1000, 2000, 5000, 10000, 20000, 30000))
	histogramOperationTableDuration, _ := meter.Float64Histogram("instruction_table_duration_seconds",
		api.WithDescription("Instruction table duration in seconds"),
		api.WithExplicitBucketBoundaries(5, 10, 15, 20, 30, 50, 100, 200, 300, 500, 1000, 2000, 5000, 10000, 20000, 30000))

	return &PrometheusAdapter{
		counterConsumerContext:           counterConsumerContext,
		histogramConsumerContextDuration: histogramConsumerContextDuration,
		histogramOperationTableDuration:  histogramOperationTableDuration,
	}
}
func (p *PrometheusAdapter) CounterConsumerContext(ctx context.Context, context, result string) {
	opt := api.WithAttributes(
		attribute.Key("context").String(context),
		attribute.Key("result").String(result),
	)
	p.counterConsumerContext.Add(ctx, 1, opt)
}
func (p *PrometheusAdapter) HistogramConsumerContextDuration(ctx context.Context, context, result string, duration float64) {
	opt := api.WithAttributes(
		attribute.Key("context").String(context),
		attribute.Key("result").String(result),
	)
	p.histogramConsumerContextDuration.Record(ctx, duration, opt)
}
func (p *PrometheusAdapter) HistogramOperationTableDuration(ctx context.Context, database, table, method string, duration float64) {
	opt := api.WithAttributes(
		attribute.Key("database").String(database),
		attribute.Key("table").String(table),
		attribute.Key("method").String(method),
	)
	p.histogramOperationTableDuration.Record(ctx, duration, opt)
}
