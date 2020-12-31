package configs

import (
	"os"
)

const (
	// Default config definitions
	defProcName     = "iot-api"
	defHTTPBindAddr = "0.0.0.0:5000"
	defGRPCBindAddr = "0.0.0.0:5010"

	// Environment variable names
	envProcName     = "PROC_NAME"
	envHTTPBindAddr = "HTTP_BIND_ADDR"
	envGRPCBindAddr = "GRPC_BIND_ADDR"
)

// Config config struct definition
type Config struct {
	ProcessName  string
	GRPCBindAddr string
	HTTPBindAddr string
}

// env get environment variable or fallback to default one
func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// LoadConfig load config parameters
func LoadConfig() *Config {
	return &Config{
		ProcessName:  env(envProcName, defProcName),
		GRPCBindAddr: env(envGRPCBindAddr, defGRPCBindAddr),
		HTTPBindAddr: env(envHTTPBindAddr, defHTTPBindAddr),
	}
}
