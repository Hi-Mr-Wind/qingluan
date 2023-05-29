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
}

// SetCapacity 设置队列容量
func SetCapacity(q *queueImpl.StandardQueue, capacity int32) {
	q.Capacity = capacity
}

// SetPatternCopy 设置队列模式 true为复制模式  否则为分发模式
func SetPatternCopy(q *queueImpl.StandardQueue, patternCopy bool) {
	q.PatternCopy = patternCopy
}
