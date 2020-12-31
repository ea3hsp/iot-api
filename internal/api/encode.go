package api

import (
	"context"

	"github.com/ea3hsp/iot-api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// EncodeGRPCPostMsg ...
func EncodeGRPCPostMsg(ctx context.Context, r interface{}) (interface{}, error) {
	return &pb.PostMsgResp{
		Msg:       "",
		Timestamp: &timestamppb.Timestamp{},
	}, nil
}
