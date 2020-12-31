package configs

import (
	"os"
)

const (
	// Default config definitions
	defProcName   = "hermes-device-api"
	defDeviceAddr = "0.0.0.0:3000"
	defEventAddr  = "nats://10.6.4.94:4222/"
	defMqttAddr   = "tcp://10.6.4.94:19000"
	defMqttUser   = "admin"
	defMqttPass   = "Me8140@01"
	defQueueAddr  = "10.6.4.94:11300"
	defNumWorkers = "2"

	// Environment variable names
	envProcName   = "PROCESS_NAME"
	envDeviceAddr = "DEVICE_API_ADDRESS"
	envEventAddr  = "EVENT_BROKER_ADDRESS"
	envMqttAddr   = "MQTT_BROKER_ADDRESS"
	envMqttUser   = "MQTT_USER"
	envMqttPass   = "MQTT_PASS"
	envQueueAddr  = "QUEUE_ADDRESS"
	envNumWorkers = "NUM_WORKERS"
)

// Config config struct definition
type Config struct {
	ProcessName string
	DeviceAddr  string
	EventAddr   string
	MqttAddr    string
	MqttUser    string
	MqttPass    string
	QueueAddr   string
	NumWorkers  string
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
		ProcessName: env(envProcName, defProcName),
		DeviceAddr:  env(envDeviceAddr, defDeviceAddr),
		EventAddr:   env(envEventAddr, defEventAddr),
		MqttAddr:    env(envMqttAddr, defMqttAddr),
		MqttUser:    env(envMqttUser, defMqttUser),
		MqttPass:    env(envMqttPass, defMqttPass),
		QueueAddr:   env(envQueueAddr, defQueueAddr),
		NumWorkers:  env(envNumWorkers, defNumWorkers),
	}
}
