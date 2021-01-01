package models

import "time"

// PostTelemetryReq post message request
type PostTelemetryReq struct {
	DeviceID  string    `json:"deviceid"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

// PostTelemetryResp post message response
type PostTelemetryResp struct {
	Msg       string    `json:"msg"`
	Timestamp time.Time `json:"timestamp"`
}
