package message

import "sync"

type Message interface {
	GetHeader() *Heard
}

// Heard 消息头
type Heard struct {
	//消息id
	MessageId string `json:"id"`
	//生产者ID
	ProducerId string `json:"producerID"`
	//所属队列ID
	QueueId string `json:"queueID"`
	// 消息生成的时间戳
	Timestamp int64 `json:"timestamp"`
	//发送时间
	SendTime int64 `json:"sendTime"`
	//有效期
	Indate int64 `json:"indate"`
}

// SubMessage 订阅消息
type SubMessage struct {
	//消息锁
	lock sync.Mutex
	//消息头
	Heard `json:"header"`
	//订阅者id
	Subscriber []string `json:"subscriber"`
	//正文
	Body []byte `json:"body"`
}

func (message *SubMessage) GetHeader() *Heard {
	return &message.Heard
}

// RandomMessage 随机消息
type RandomMessage struct {
	//消息锁
	lock sync.Mutex
	//消息头
	Heard `json:"heard"`
	//随机权重
	Weight int32 `json:"weight,omitempty"`
	//可消费次数
	Number int32 `json:"number,omitempty"`
	//消息正文
	Body []byte `json:"body,omitempty"`
}

func (message *RandomMessage) GetHeader() *Heard {
	return &message.Heard
}

// ExclusiveMessage 独享消息
// 独享消息只能有一个队列消费一次，消费后则消息
type ExclusiveMessage struct {
	//消息锁
	lock sync.Mutex
	//消息头
	Heard `json:"heard"`
	//消息正文
	Body []byte `json:"body,omitempty"`
}

func (message *ExclusiveMessage) GetHeader() *Heard {
	return &message.Heard
}

// WeightMessage 权重消息
type WeightMessage struct {
	//消息锁
	lock sync.Mutex
	//消息头
	Heard `json:"heard"`
	//消息权重
	Weight int32 `json:"weight,omitempty"`
	//消息正文
	Body []byte `json:"body,omitempty"`
}

func (message *WeightMessage) GetHeader() *Heard {
	return &message.Heard
}

// DelayedMessage 延迟消息
type DelayedMessage struct {
	//消息锁
	lock sync.Mutex
	//消息头
	Heard `json:"heard"`
	// 延迟时间 毫秒
	Delayed int64 `json:"delayed"`
	//消息正文
	Body []byte `json:"body,omitempty"`
}

func (message *DelayedMessage) GetHeader() *Heard {
	return &message.Heard
}
