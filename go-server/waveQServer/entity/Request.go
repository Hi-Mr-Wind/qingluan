package entity

import "time"

// Request 拉取消息请求信息
type Request struct {
	//请求时间
	ResTime time.Time
	//消费者ID
	ApiKey string
	//前条消息ID
	FormerId string
}

// NewRequest 获取消息拉取请求实体
func NewRequest() *Request {
	return new(Request)
}
