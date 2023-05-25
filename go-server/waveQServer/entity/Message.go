package entity

import "time"

// Heard 消息头
type Heard struct {
	//消息id
	Id []byte `json:"id"`
	//生产者ID
	ProducerID []byte `json:"producerID"`
	//发送时间
	SendTime time.Time `json:"sendTime"`
	//所属队列ID
	QueueID []byte `json:"queueID"`
	//前条消息ID
	FormerId []byte `json:"formerId"`
	//有效期
	Indate int32 `json:"indate"`
	//延迟时间（毫秒）
	DelayTime uint32 `json:"delayTime"`
	//消息状态
	State int8 `json:"state"`
}

// Message 消息对象
type Message struct {
	Header Heard `json:"header"`

	Body []byte `json:"body"`
}
