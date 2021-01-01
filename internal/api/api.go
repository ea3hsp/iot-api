package api

import (
	"context"
	"time"

	"github.com/ea3hsp/iot-api/internal/backend"
	"github.com/ea3hsp/iot-api/internal/models"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger  log.Logger
	backend backend.DomoBackend
}

// NewService creates svc
func NewService(backend backend.DomoBackend, logger log.Logger) DomoService {
	return &service{
		backend: backend,
		logger:  logger,
	}
}

func (svc *service) PostTelemetry(ctx context.Context, req models.PostTelemetryReq) (models.PostTelemetryResp, error) {
	// creates logger
	logger := log.With(svc.logger, "method", "PostMsg")
	level.Info(logger).Log("msg", "posting message")
	return models.PostTelemetryResp{
		Msg:       "ok",
		Timestamp: time.Now().UTC(),
	}, nil
}
