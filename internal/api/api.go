package api

import (
	"context"
	"time"

	"github.com/ea3hsp/iot-api/internal/models"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger log.Logger
}

// NewService creates svc
func NewService(logger log.Logger) DomoService {
	return &service{
		logger: logger,
	}
}

func (svc *service) PostMsg(ctx context.Context, req models.PostMsgReq) (models.PostMsgResp, error) {
	// creates logger
	logger := log.With(svc.logger, "method", "PostMsg")
	level.Info(logger).Log("msg", "posting message")
	return models.PostMsgResp{
		Msg:       "ok",
		Timestamp: time.Now().UTC(),
	}, nil
}
