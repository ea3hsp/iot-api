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
	"github.com/ea3hsp/iot-api/internal/backend/mqtt"

	httptransport "github.com/ea3hsp/iot-api/internal/api/http"
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
	pidFile := fmt.Sprintf("%s.pid", cfg.ProcessName)
	upg, err := tableflip.New(tableflip.Options{
		PIDFile: pidFile,
	})
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
	// //
	// errs := make(chan error)
	// go func() {
	// 	c := make(chan os.Signal)
	// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// 	errs <- fmt.Errorf("%s", <-c)
	// }()
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
	// creates device endpoints
	endpoints := api.MakeEndpoints(srv)
	// services context
	ctx := context.Background()

	// // creates REST API Server
	// go func() {
	// 	// banner
	// 	level.Info(logger).Log("msg", fmt.Sprintf("domo api worker API listening: %s", cfg.HTTPBindAddr))
	// 	// service http handler
	// 	hdl := httptransport.NewHTTPServer(ctx, endpoints)
	// 	// start http server
	// 	http.ListenAndServe(cfg.HTTPBindAddr, hdl)
	// }()

	// creates REST API Server
	// Listen must be called before Ready
	ln, err := upg.Listen("tcp", cfg.HTTPBindAddr)
	if err != nil {
		level.Error(logger).Log("msg", fmt.Sprintf("Can't listen: %s", err.Error()))
		os.Exit(1)
	}
	// service http handler
	hdl := httptransport.NewHTTPServer(ctx, endpoints)
	// set server params
	server := http.Server{
		Handler:           hdl,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
	}
	//
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
