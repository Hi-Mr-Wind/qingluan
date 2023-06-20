package message

import "sync"

type Message interface {
	GetHeader() *Heard
}

// Heard 消息头
type Heard struct {
	//消息id
	Id string `json:"id"`
	//生产者ID
	ProducerID string `json:"producerID"`
	// 消息生成的时间戳
	Timestamp int64 `json:"timestamp"`
	//发送时间
	SendTime int64 `json:"sendTime"`
	//所属队列ID
	QueueID string `json:"queueID"`
	//有效期
	Indate int32 `json:"indate"`
}

// SubMessage 订阅消息
type SubMessage struct {
	//消息锁
	lock sync.Mutex
	//消息头
	Header Heard `json:"header"`
	//订阅者id
	Subscriber []string `json:"subscriber"`
	//正文
	Body []byte `json:"body"`
}

func (message *SubMessage) GetHeader() *Heard {
	return &message.Header
}

// RandomMessage 随机消息
type RandomMessage struct {
	//消息锁
	lock sync.Mutex
	//消息头
	Heard Heard `json:"heard"`
	//随机权重
	Weight int `json:"weight,omitempty"`
	//可消费次数
	Number int `json:"number,omitempty"`
	//消息正文
	Body []byte `json:"body,omitempty"`
}

func (message *RandomMessage) GetHeader() *Heard {
	return &message.Heard
}

// ExclusiveMessage 独享消息
// 独享消息只能有一个队列消费一次，消费后则消息
type ExclusiveMessage struct {
}
