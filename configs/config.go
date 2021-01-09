package configs

import (
	"os"
)

const (
	// Default config definitions
	defProcName     = "iot-api-worker"
	defHTTPBindAddr = "0.0.0.0:5000"
	defMqttAddr     = "tcp://localhost:1883"
	defMqttUser     = ""
	defMqttPass     = ""

	// Environment variable names
	envProcName     = "PROC_NAME"
	envHTTPBindAddr = "HTTP_BIND_ADDR"
	envMqttAddr     = "MQTT_ADDRESS"
	envMqttUser     = "MQTT_USER"
	envMqttPass     = "MQTT_PASS"
)

// Config config struct definition
type Config struct {
	ProcessName    string
	HTTPBindAddr   string
	MQTTBrokerAddr string
	MQTTUser       string
	MQTTPass       string
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
		ProcessName:    env(envProcName, defProcName),
		HTTPBindAddr:   env(envHTTPBindAddr, defHTTPBindAddr),
		MQTTBrokerAddr: env(envMqttAddr, defMqttAddr),
		MQTTUser:       env(envMqttUser, defMqttUser),
		MQTTPass:       env(envMqttPass, defMqttPass),
	}
}
