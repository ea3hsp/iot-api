package api

import (
	"context"

	"github.com/ea3hsp/iot-api/internal/models"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints endpoints definions
type Endpoints struct {
	CreatePostMsg endpoint.Endpoint
}

// MakeEndpoints creates enpoints logic
func MakeEndpoints(svc DomoService) Endpoints {
	return Endpoints{
		CreatePostMsg: func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(models.PostMsgReq)
			return svc.PostMsg(ctx, req)
		},
	}
}
