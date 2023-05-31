package queue

import (
	"waveQServer/core/queue/queueImpl"
	"waveQServer/entity"
	"waveQServer/identity"
)

type Queue interface {
	Size() int32
	Push(message *entity.Message)
	AddUser(user *identity.User)
	GetQueueId() string
}

// SetCapacity 设置队列容量
func SetCapacity(q *queueImpl.StandardQueue, capacity int32) {
	q.Capacity = capacity
}
