package entity

import "time"

// Heard 消息头
type Heard struct {
	//消息id
	Id []byte
	//生产者ID
	ProducerID []byte
	//发送时间
	SendTime time.Time
	//所属队列ID
	QueueID []byte
	//前条消息ID
	FormerId []byte
	//有效期
	Indate int32
	//延迟时间（毫秒）
	DelayTime uint32
	//消息状态
	State int8
}

// Message 消息对象
type Message struct {
	Header Heard

	Body []byte
}
