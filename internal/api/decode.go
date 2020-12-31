package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ea3hsp/iot-api/internal/models"
)

// DecodePostMsg decodes device msg device
func DecodePostMsg(ctx context.Context, r *http.Request) (interface{}, error) {
	var res models.PostMsgResp
	json.NewDecoder(r.Body).Decode(&res)
	return res, nil
}
