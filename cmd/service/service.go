package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ea3hsp/iot-api/configs"
	"github.com/ea3hsp/iot-api/internal/api"
	grpctransport "github.com/ea3hsp/iot-api/internal/api/grpc"
	httptransport "github.com/ea3hsp/iot-api/internal/api/http"
	"github.com/ea3hsp/iot-api/pb"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunService execute main function
func RunService() {
	// parse os args
	cfg := configs.LoadConfig()
	// Creates logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "hermes-device-api",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	//
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	// creates service
	var srv api.DomoService
	{
		srv = api.NewService(logger)
	}
	// loggging middleware attachment
	srv = api.LoggingMiddleware(srv, logger)
	srv = api.MetricsMiddleware(
		srv,
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "device",
			Subsystem: "api",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "device",
			Subsystem: "api",
			Name:      "request_latency_seconds",
			Help:      "Total duration of requests in seconds.",
		}, []string{"method"}),
	)
	// creates device endpoints
	endpoints := api.MakeEndpoints(srv)
	// services context
	ctx := context.Background()
	// creates REST API Server
	go func() {
		// banner
		level.Info(logger).Log("msg", fmt.Sprintf("hermes device API listening: %s", cfg.HTTPBindAddr))
		// service http handler
		hdl := httptransport.NewHTTPServer(ctx, endpoints)
		// start http server
		http.ListenAndServe(cfg.HTTPBindAddr, hdl)
	}()
	// creates GRPC API Server
	go func() {
		// banner
		level.Info(logger).Log("msg", fmt.Sprintf("hermes device GRPC API listening: %s", cfg.GRPCBindAddr))
		// creates tcp channel comunication
		listener, err := net.Listen("tcp", cfg.GRPCBindAddr)
		if err != nil {
			errs <- err
			return
		}
		// grpc server instance
		gRPCServer := grpc.NewServer()
		// register of grpc server reflection
		reflection.Register(gRPCServer)
		// registers rpc services
		pb.RegisterDeviceServiceServer(gRPCServer, grpctransport.NewGRPCServer(ctx, endpoints, logger))
		// start service
		errs <- gRPCServer.Serve(listener)
	}()
	level.Error(logger).Log("msg", fmt.Sprintf("exit %v", <-errs))
}
