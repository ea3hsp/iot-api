package grpc

import (
	"context"

	"github.com/ea3hsp/iot-api/internal/api"
	"github.com/ea3hsp/iot-api/pb"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	postMsg grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC DeviceService.
func NewGRPCServer(ctx context.Context, endpoints api.Endpoints, logger log.Logger) pb.DomoServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	return &grpcServer{
		postMsg: grpctransport.NewServer(
			endpoints.CreatePostMsg,
			api.DecodeGRPCPostMsg,
			api.EncodeGRPCPostMsg,
			options...,
		),
	}
}

func (s *grpcServer) PostMsg(ctx context.Context, req *pb.PostMsgReq) (*pb.PostMsgResp, error) {
	_, resp, err := s.postMsg.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PostMsgResp), nil
}
