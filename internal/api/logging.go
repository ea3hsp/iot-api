package api

import (
	"context"
	"fmt"
	"time"

	"github.com/ea3hsp/iot-api/internal/models"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type loggingMiddleware struct {
	logger log.Logger
	svc    DomoService
}

// LoggingMiddleware adds logging facilities to the adapter.
func LoggingMiddleware(svc DomoService, logger log.Logger) DomoService {
	return &loggingMiddleware{logger, svc}
}

func (lm *loggingMiddleware) PostTelemetry(ctx context.Context, req models.PostTelemetryReq) (res models.PostTelemetryResp, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method PostTelemetry took %s to complete", time.Since(begin))
		if err != nil {
			level.Error(lm.logger).Log("msg", fmt.Sprintf("%s with error: %s.", message, err.Error()))
			return
		}
		level.Info(lm.logger).Log("msg", fmt.Sprintf("%s without errors.", message))
	}(time.Now())
	return lm.svc.PostTelemetry(ctx, req)
}
