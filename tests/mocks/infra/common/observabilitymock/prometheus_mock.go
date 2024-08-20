package observabilitymock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type PrometheusAdapterMock struct {
	mock.Mock
}

func (p *PrometheusAdapterMock) CounterRequestHttpStatusCode(ctx context.Context, router string, statusCode int) {
	p.Called(ctx, router, statusCode)
}
func (p *PrometheusAdapterMock) CounterRedisError(ctx context.Context, method string, codeError string) {
	p.Called(ctx, method, codeError)
}
func (p *PrometheusAdapterMock) HistogramRequestDuration(ctx context.Context, router string, statusCode int, duration float64) {
	p.Called(ctx, router, statusCode, duration)
}
func (p *PrometheusAdapterMock) HistogramInstructionTableDuration(ctx context.Context, database, table, method string, duration float64) {
	p.Called(ctx, database, table, method, duration)
}

func (p *PrometheusAdapterMock) HistogramOperationDuration(ctx context.Context, operation string, context string, duration float64) {
	p.Called(ctx, operation, context, duration)
}
