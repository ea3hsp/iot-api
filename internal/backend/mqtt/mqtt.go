package mqtt

import (
	"crypto/tls"
	"fmt"
	"hash/fnv"
	"sync"
	"time"

	"github.com/ea3hsp/iot-api/internal/models"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var (
	// PublishQoS Quality of Service Level
	PublishQoS byte = 0x00
	// SubscribeQoS Quality of Service Level
	SubscribeQoS byte = 0x00
	// PublishTimeout is the timeout before returning from publish without checking error
	PublishTimeout = 50 * time.Millisecond
	// BufferSize indicates the maximum number of MQTT messages that should be buffered
	BufferSize = 10
	// ConnectRetries says how many times the client should retry a failed connection
	ConnectRetries = 10
	// ConnectRetryDelay says how long the client should wait between retries
	ConnectRetryDelay = time.Second
)

const (
	// TelemetryTopicFormat telemetry topic format
	TelemetryTopicFormat = "domo/devices/%s/telemetry"
	// StatsTopicFormat statistics topic format
	StatsTopicFormat = "domo/devices/%s/attributes/stats"
	// StateTopicFormat state topic format
	StateTopicFormat = "domo/devices/%s/attributes/state"
	// InfoTopicFormat info topic format
	InfoTopicFormat = "domo/devices/%s/attributes/info"
	// LogTopicFormat log topic format
	LogTopicFormat = "domo/devices/%s/log"
	// DisconnectTopicFormat disconnect topic format
	DisconnectTopicFormat = "domo/devices/%s/disconnect"
)

// Config contains configuration for MQTT
type Config struct {
	Brokers   []string
	Username  string
	Password  string
	TLSConfig *tls.Config
}

type subscription struct {
	handler paho.MessageHandler
	cancel  func()
}

// MQTT side of the bridge
type MQTT struct {
	logger        log.Logger
	client        paho.Client
	subscriptions map[string]subscription
	mu            sync.Mutex
	cfg           Config
}

func clientID() string {
	t := time.Now().String()
	h := fnv.New32a()
	h.Write([]byte(t))
	r := fmt.Sprintf("%x", h.Sum32())
	return r
}

// New returns a new MQTT
func New(config Config, logger log.Logger) (*MQTT, error) {
	mqtt := new(MQTT)
	mqtt.logger = logger
	mqtt.cfg = config
	mqttOpts := paho.NewClientOptions()
	for _, broker := range config.Brokers {
		mqttOpts.AddBroker(broker)
	}
	if config.TLSConfig != nil {
		mqttOpts.SetTLSConfig(config.TLSConfig)
	}
	mqttOpts.SetClientID(fmt.Sprintf("domo-api-worker-%s", clientID()))

	mqttOpts.SetUsername(config.Username)
	mqttOpts.SetPassword(config.Password)

	mqttOpts.SetKeepAlive(30 * time.Second)
	mqttOpts.SetPingTimeout(10 * time.Second)
	mqttOpts.SetCleanSession(true)
	mqttOpts.SetDefaultPublishHandler(func(_ paho.Client, msg paho.Message) {
		level.Warn(mqtt.logger).Log("msg", fmt.Sprintf("Received unhandled message on MQTT: %v", msg))
	})

	mqtt.subscriptions = make(map[string]subscription)
	var reconnecting bool
	mqttOpts.SetConnectionLostHandler(func(_ paho.Client, err error) {
		level.Warn(mqtt.logger).Log("msg", fmt.Sprintf("Disconnected (%s). Reconnecting...", err.Error()))
		reconnecting = true
	})
	mqttOpts.SetOnConnectHandler(func(_ paho.Client) {
		if reconnecting {
			mqtt.resubscribe()
			reconnecting = false
		}
	})
	mqtt.client = paho.NewClient(mqttOpts)

	return mqtt, nil
}

// Connect to MQTT
func (c *MQTT) Connect() error {
	var err error
	for retries := 0; retries < ConnectRetries; retries++ {
		token := c.client.Connect()
		finished := token.WaitTimeout(1 * time.Second)
		if !finished {
			level.Warn(c.logger).Log("msg", "MQTT connection took longer than expected...")
			token.Wait()
		}
		err = token.Error()
		if err == nil {
			break
		}
		level.Warn(c.logger).Log("msg", fmt.Sprintf("Could not connect to MQTT (%s). Retrying...", err.Error()))
		<-time.After(ConnectRetryDelay)
	}
	if err != nil {
		return fmt.Errorf("Could not connect to MQTT (%s)", err)
	}

	return err
}

// Disconnect from MQTT
func (c *MQTT) Disconnect() error {
	level.Warn(c.logger).Log("msg", "mqtt broker disconnecting")
	c.client.Disconnect(100)
	return nil
}

func (c *MQTT) resubscribe() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for topic, subscription := range c.subscriptions {
		c.client.Subscribe(topic, SubscribeQoS, subscription.handler)
	}
}

func (c *MQTT) publish(topic string, payload []byte) paho.Token {
	return c.client.Publish(topic, PublishQoS, false, payload)
}

// PublishTelemetry publish telemetry message
func (c *MQTT) PublishTelemetry(msg *models.PostTelemetryReq) error {
	topic := fmt.Sprintf(TelemetryTopicFormat, msg.DeviceID)
	token := c.publish(topic, []byte(msg.Payload))
	go func() {
		token.Wait()
		if err := token.Error(); err != nil {
			level.Warn(c.logger).Log("msg", fmt.Sprintf("could not publish telemetry message topic=%s device=%s err=%s", topic, msg.DeviceID, err.Error()))
			return
		}
		level.Info(c.logger).Log("msg", "published telemetry message")
	}()
	return nil
}
