package backend

// DomoBackend domo backend implementation
type DomoBackend interface {
	Connect() error
	Disconnect() error
	PublishTelemetry(msg []byte) error
}
