package observability

import "context"

type IMetricAdapter interface {
	CounterRequestHttpStatusCode(ctx context.Context, router string, statusCode int)
	HistogramRequestDuration(ctx context.Context, router string, statusCode int, duration float64)
	HistogramInstructionTableDuration(ctx context.Context, database, table, method string, duration float64)
}
