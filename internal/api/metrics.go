package api

import (
	"context"
	"time"

	"github.com/ea3hsp/iot-api/internal/models"
	"github.com/go-kit/kit/metrics"
)

type metricsMiddleware struct {
	counter metrics.Counter
	latency metrics.Histogram
	svc     DomoService
}

// MetricsMiddleware instruments adapter by tracking request count and latency.
func MetricsMiddleware(svc DomoService, counter metrics.Counter, latency metrics.Histogram) DomoService {
	return &metricsMiddleware{
		counter: counter,
		latency: latency,
		svc:     svc,
	}
}

func (mm *metricsMiddleware) PostTelemetry(ctx context.Context, req models.PostTelemetryReq) (models.PostTelemetryResp, error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "PostTelemetry").Add(1)
		mm.latency.With("method", "PostTelemetry").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mm.svc.PostTelemetry(ctx, req)
}
