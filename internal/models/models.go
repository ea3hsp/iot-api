package models

import "time"

// PostMsgReq post message request
type PostMsgReq struct {
	DeviceID  string    `json:"deviceid"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

// PostMsgResp post message response
type PostMsgResp struct {
	Msg       string    `json:"msg"`
	Timestamp time.Time `json:"timestamp"`
}
