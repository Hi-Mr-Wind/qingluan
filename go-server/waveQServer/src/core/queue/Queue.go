package queue

import (
	"waveQServer/src/core/message"
	"waveQServer/src/core/queue/queueImpl"
	"waveQServer/src/entity"
)

type Queue interface {
	Size() int32
	Push(message *message.Message)
	AddUser(user *entity.User)
	GetQueueId() string
}

// SetCapacity 设置队列容量
func SetCapacity(q *queueImpl.BroadcastQueue, capacity int32) {
	q.Capacity = capacity
}
