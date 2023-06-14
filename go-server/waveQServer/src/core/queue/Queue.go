package queue

import (
	"waveQServer/src/core/queue/queueImpl"
	"waveQServer/src/entity/message"
	"waveQServer/src/identity"
)

type Queue interface {
	Size() int32
	Push(message *message.Message)
	AddUser(user *identity.User)
	GetQueueId() string
}

// SetCapacity 设置队列容量
func SetCapacity(q *queueImpl.StandardQueue, capacity int32) {
	q.Capacity = capacity
}