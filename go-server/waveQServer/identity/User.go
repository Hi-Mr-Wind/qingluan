package identity

import (
	"waveQServer/entity"
)

// Subscriber 定义一个消息通道
type subscriber chan *entity.Message

// User 消费者用户结构
type User struct {
	//apikey
	ApiKey string `json:"apiKey"`
	//访问队列权限
	RccessRights [][]byte `json:"rccessRights"`
	//过期时间
	OutTime int64
	//消费通道
	MessageChan subscriber
	//消费游标
	Index int32
}

type UserBuilder struct {
	rccessRights [][]byte
	apiKey       string
	outTime      int64
}

// BuilderUser 获取建造者对象
func BuilderUser() *UserBuilder {
	return new(UserBuilder)
}

// addRccessRights 添加访问权限组
func (b *UserBuilder) addRccessRights(rccessRights [][]byte) *UserBuilder {
	b.rccessRights = rccessRights
	return b
}

// 添加apikey
func (b *UserBuilder) addApiKey(apiKey string) *UserBuilder {
	b.apiKey = apiKey
	return b
}

// 添加过期时间
func (b *UserBuilder) addOutTime(outTime int64) *UserBuilder {
	b.outTime = outTime
	return b
}

// 构建user对象
func (b *UserBuilder) build() *User {
	user := new(User)
	user.MessageChan = make(subscriber)
	user.Index = 0
	user.RccessRights = b.rccessRights
	user.ApiKey = b.apiKey
	user.RccessRights = b.rccessRights
	user.OutTime = b.outTime
	return user
}
