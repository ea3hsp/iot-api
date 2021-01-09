package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ea3hsp/iot-api/configs"
	"github.com/ea3hsp/iot-api/internal/api"
	httptransport "github.com/ea3hsp/iot-api/internal/api/http"
	"github.com/ea3hsp/iot-api/internal/backend/mqtt"
	basicauth "github.com/go-kit/kit/auth/basic"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/cloudflare/tableflip"
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
			"service", cfg.ProcessName,
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	upg, err := tableflip.New(tableflip.Options{})
	if err != nil {
		panic(err)
	}
	defer upg.Stop()
	// Do an upgrade on SIGHUP
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)
		for range sig {
			err := upg.Upgrade()
			if err != nil {
				level.Error(logger).Log("msg", fmt.Sprintf("Upgrade failed: %s", err.Error()))
			}
		}
	}()
	// create mqtt config
	mqttCfg := mqtt.Config{
		Brokers: []string{
			cfg.MQTTBrokerAddr,
		},
		Username:  cfg.MQTTUser,
		Password:  cfg.MQTTPass,
		TLSConfig: nil,
	}
	// create mqtt backend
	backend, err := mqtt.New(mqttCfg, logger)
	if err != nil {
		level.Error(logger).Log("msg", fmt.Sprintf("create backend error=%s", err.Error()))
		os.Exit(1)
	}
	// creates service
	var srv api.DomoService
	{
		srv = api.NewService(backend, logger)
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
	var middlewares []endpoint.Middleware
	auth := basicauth.AuthMiddleware("admin", "hermes", "myrealm")
	middlewares = append(middlewares, auth)
	// creates device endpoints
	endpoints := api.MakeEndpoints(srv, middlewares)
	// services context
	ctx := context.Background()
	// creates REST API Server
	// Listen must be called before Ready
	ln, err := upg.Listen("tcp", cfg.HTTPBindAddr)
	if err != nil {
		level.Error(logger).Log("msg", fmt.Sprintf("Can't listen: %s", err.Error()))
		os.Exit(1)
	}
	// banner
	level.Info(logger).Log("msg", fmt.Sprintf("domo api worker API listening: %s", cfg.HTTPBindAddr))
	// service http handler
	hdl := httptransport.NewHTTPServer(ctx, endpoints)
	// set server params
	server := http.Server{
		Handler:           hdl,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	// http server go routine
	go func() {
		// start http server
		server.Serve(ln)
		if err != http.ErrServerClosed {
			level.Error(logger).Log("msg", fmt.Sprintf("Can't listen: %s", err.Error()))
		}
	}()
	//
	if err := upg.Ready(); err != nil {
		panic(err)
	}
	<-upg.Exit()
	//
	time.AfterFunc(30*time.Second, func() {
		level.Info(logger).Log("msg", "Graceful shutdown timed out")
		os.Exit(1)
	})
	// Wait for connections to drain.
	server.Shutdown(context.Background())
	//level.Error(logger).Log("msg", fmt.Sprintf("exit %v", <-errs))
	backend.Disconnect()
}
