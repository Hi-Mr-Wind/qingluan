package identity

import (
	"waveQServer/src/core/message"
)

// Subscriber 定义一个消息通道
type subscriber chan *message.Message

// User 消费者用户结构
type User struct {
	//apikey
	ApiKey string `json:"apiKey"`
	//访问队列权限
	RccessRights []string `json:"rccessRights"`
	//过期时间
	OutTime int64
	//消费通道
	MessageChan subscriber
	//消费游标
	Index int32

	GroupId string

	QueueId string
}

type UserBuilder struct {
	rccessRights []string
	apiKey       string
	outTime      int64
}

// BuilderUser 获取建造者对象
func BuilderUser() *UserBuilder {
	return new(UserBuilder)
}

// AddRccessRights 添加访问权限组
func (b *UserBuilder) AddRccessRights(rccessRights []string) *UserBuilder {
	b.rccessRights = rccessRights
	return b
}

// AddApiKey 添加apikey
func (b *UserBuilder) AddApiKey(apiKey string) *UserBuilder {
	b.apiKey = apiKey
	return b
}

// AddOutTime 添加过期时间
func (b *UserBuilder) AddOutTime(outTime int64) *UserBuilder {
	b.outTime = outTime
	return b
}

// Build 构建user对象
func (b *UserBuilder) Build() *User {
	user := new(User)
	user.MessageChan = make(subscriber)
	user.Index = 0
	user.RccessRights = b.rccessRights
	user.ApiKey = b.apiKey
	user.RccessRights = b.rccessRights
	user.OutTime = b.outTime
	return user
}
