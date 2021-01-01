package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ea3hsp/iot-api/internal/models"
)

// DecodeTelemetry decodes http rest api post messages
func DecodeTelemetry(ctx context.Context, r *http.Request) (interface{}, error) {
	var res models.PostTelemetryReq
	json.NewDecoder(r.Body).Decode(&res)
	return res, nil
}
