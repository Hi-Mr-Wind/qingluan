package entity

// NewAdmin 创建admin对象
func NewAdmin() *Admin {
	return new(Admin)
}

// NewUser 创建User对象
func NewUser() *User {
	return new(User)
}

// NewGroups 创建消息组对象
func NewGroups() *Groups {
	return new(Groups)
}

// NewQueue 创建队列对象
func NewQueue() *Queue {
	return new(Queue)
}

// NewQueueUser 创建队列用户对应
func NewQueueUser() *QueueUser {
	return new(QueueUser)
}

// NewGroupQueue 组-队列对应
func NewGroupQueue() *GroupQueue {
	return new(GroupQueue)
}
