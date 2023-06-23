package entity

// User 消费者用户
type User struct {
	Id             string `gorm:"column:id;primary_key;NOT NULL"`
	ApiKey         string `gorm:"column:api_key"`
	ExpirationTime int64  `gorm:"column:expiration_time"`
	CreatTime      string `gorm:"column:creat_time"`
}

func (u *User) TableName() string {
	return "user"
}

// Admin 管理员用户
type Admin struct {
	Id        string `gorm:"column:id;primary_key;NOT NULL"`
	UserName  string `gorm:"column:user_name"`
	Password  string `gorm:"column:password"`
	CreatTime string `gorm:"column:creat_time"`
}

func (a *Admin) TableName() string {
	return "admin"
}

// Groups 队列组
type Groups struct {
	GroupId   string `gorm:"column:group_id;primary_key;NOT NULL"`
	CreatTime string `gorm:"column:creat_time"`
}

func (g *Groups) TableName() string {
	return "groups"
}

// Queue 队列
type Queue struct {
	QueueId   string `gorm:"column:queue_id;primary_key;NOT NULL"`
	QueueType string `gorm:"column:queue_type"`
	capacity  int32  `gorm:"column:capacity"`
	CreatTime string `gorm:"column:creat_time"`
}

func (q *Queue) TableName() string {
	return "queue"
}

// QueueUser 队列用户对应
type QueueUser struct {
	Id      string `gorm:"column:id;primary_key;NOT NULL"`
	UserId  string `gorm:"column:user_id"`
	QueueId string `gorm:"column:queue_id"`
}

func (q *QueueUser) TableName() string {
	return "queue_ueser"
}

// GroupQueue 组队列对应
type GroupQueue struct {
	Id      string `gorm:"column:id;primary_key;NOT NULL"`
	GroupId string `gorm:"column:group_id"`
	QueueId string `gorm:"column:queue_id"`
}

func (g *GroupQueue) TableName() string {
	return "group_queue"
}

// RandomMessage 权重随机消息
type RandomMessage struct {
	//消息ID
	MessageId string `gorm:"column:message_id;primary_key"`
	// 生产者ID
	ProducerId string `gorm:"column:producer_id"`
	// 消息生成时间
	Timestamp int64 `gorm:"column:timestamp"`
	// 消息发送时间
	SendTime int64 `gorm:"column:send_time"`
	//所属队列ID
	QueueId string `gorm:"column:queue_id"`
	//有效期
	Indate int64 `gorm:"column:indate"`
	//权重
	Weight int32 `gorm:"column:weight"`
	//可消费次数
	Number int32 `gorm:"column:number"`
	//消息正文
	Body []byte `gorm:"column:body"`
}

func (m *RandomMessage) TableName() string {
	return "random_message"
}

// ExclusiveMessage 独享消息
type ExclusiveMessage struct {
	//消息ID
	MessageId string `gorm:"column:message_id;primary_key"`
	// 生产者ID
	ProducerId string `gorm:"column:producer_id"`
	// 消息生成时间
	Timestamp int64 `gorm:"column:timestamp"`
	// 消息发送时间
	SendTime int64 `gorm:"column:send_time"`
	//所属队列ID
	QueueId string `gorm:"column:queue_id"`
	//有效期
	Indate int64 `gorm:"column:indate"`
	//消息正文
	Body []byte `gorm:"column:body"`
}

func (m *ExclusiveMessage) TableName() string {
	return "exclusive_message"
}

// WeightMessage 权重消息
type WeightMessage struct {
	//消息ID
	MessageId string `gorm:"column:message_id;primary_key"`
	// 生产者ID
	ProducerId string `gorm:"column:producer_id"`
	// 消息生成时间
	Timestamp int64 `gorm:"column:timestamp"`
	// 消息发送时间
	SendTime int64 `gorm:"column:send_time"`
	//所属队列ID
	QueueId string `gorm:"column:queue_id"`
	//有效期
	Indate int64 `gorm:"column:indate"`
	//权重
	Weight int32 `gorm:"column:weight"`
	//消息正文
	Body []byte `gorm:"column:body"`
}

func (m *WeightMessage) TableName() string {
	return "weight_message"
}

// DelayedMessage 延迟消息
type DelayedMessage struct {
	//消息ID
	MessageId string `gorm:"column:message_id;primary_key"`
	// 生产者ID
	ProducerId string `gorm:"column:producer_id"`
	// 消息生成时间
	Timestamp int64 `gorm:"column:timestamp"`
	// 消息发送时间
	SendTime int64 `gorm:"column:send_time"`
	//所属队列ID
	QueueId string `gorm:"column:queue_id"`
	//有效期
	Indate int64 `gorm:"column:indate"`
	// 延迟时间 毫秒
	Delayed int64 `gorm:"column:delayed"`
	//消息正文
	Body []byte `gorm:"column:body"`
}

func (m *DelayedMessage) TableName() string {
	return "delayed_message"
}

// SubMessage 订阅消息
type SubMessage struct {
	//消息ID
	MessageId string `gorm:"column:message_id;primary_key"`
	// 生产者ID
	ProducerId string `gorm:"column:producer_id"`
	// 消息生成时间
	Timestamp int64 `gorm:"column:timestamp"`
	// 消息发送时间
	SendTime int64 `gorm:"column:send_time"`
	//所属队列ID
	QueueId string `gorm:"column:queue_id"`
	//有效期
	Indate int64 `gorm:"column:indate"`
	//订阅者id
	Subscriber string `gorm:"column:subscriber"`
	//正文
	Body []byte `gorm:"column:body"`
}

func (m *SubMessage) TableName() string {
	return "sub_message"
}
