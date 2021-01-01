package api

import (
	"context"
	"fmt"
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
	svc := new(service)
	svc.logger = logger
	svc.backend = backend
	err := svc.backend.Connect()
	if err != nil {
		level.Error(logger).Log("msg", fmt.Sprintf("new service domo error=%s", err.Error()))
		return nil
	}
	return svc
}

func (svc *service) PostTelemetry(ctx context.Context, req models.PostTelemetryReq) (models.PostTelemetryResp, error) {
	// creates logger
	logger := log.With(svc.logger, "method", "PostTelemetry")
	err := svc.backend.PublishTelemetry(&req)
	if err != nil {
		return models.PostTelemetryResp{
			Msg:       err.Error(),
			Timestamp: time.Now().UTC(),
		}, err
	}
	level.Info(logger).Log("msg", fmt.Sprintf("posting telemetry from device=%s", req.DeviceID))
	return models.PostTelemetryResp{
		Msg:       "ok",
		Timestamp: time.Now().UTC(),
	}, nil
}
