package api

import (
	"context"

	"github.com/ea3hsp/iot-api/internal/models"
)

// DomoService domo api definition
type DomoService interface {
	PostTelemetry(ctx context.Context, req models.PostTelemetryReq) (models.PostTelemetryResp, error)
}
