package api

import (
	"context"
	"time"

	"github.com/ea3hsp/iot-api/pb"
)

// EncodeGRPCPostMsg ...
func EncodeGRPCPostMsg(ctx context.Context, r interface{}) (interface{}, error) {
	return &pb.PostMsgResp{
		Msg:       "",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}, nil
}
