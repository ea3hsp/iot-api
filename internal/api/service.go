package api

import (
	"context"

	"github.com/ea3hsp/iot-api/models"
)

// DomoService domo api definition
type DomoService interface {
	PostMsg(ctx context.Context, msg *models.PostMsgReq)
}
