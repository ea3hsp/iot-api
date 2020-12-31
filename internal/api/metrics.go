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

func (mm *metricsMiddleware) PostMsg(ctx context.Context, req models.PostMsgReq) (models.PostMsgResp, error) {
	defer func(begin time.Time) {
		mm.counter.With("method", "CreateDevice").Add(1)
		mm.latency.With("method", "CreateDevice").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return mm.svc.PostMsg(ctx, req)
}
