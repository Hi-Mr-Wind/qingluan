package dto

// User 消费者用户
type User struct {
	Id             string `gorm:"column:id;primary_key;NOT NULL"`
	ApiKey         string `gorm:"column:api_key"`
	ExpirationTime string `gorm:"column:expiration_time"`
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
	CreatTime string `gorm:"column:creat_time"`
}

func (q *Queue) TableName() string {
	return "queue"
}

// QueueUeser 队列用户对应
type QueueUeser struct {
	Id      string `gorm:"column:id;primary_key;NOT NULL"`
	UserId  string `gorm:"column:user_id"`
	QueueId string `gorm:"column:queue_id"`
}

func (q *QueueUeser) TableName() string {
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
