package backend

import "github.com/ea3hsp/iot-api/internal/models"

// DomoBackend domo backend implementation
type DomoBackend interface {
	Connect() error
	Disconnect() error
	PublishTelemetry(msg *models.PostTelemetryReq) error
}
