package queue

import (
	"errors"
	"sync"
	"time"
	"waveQServer/src/comm/enum"
	"waveQServer/src/core/message"
)

// Queue 队列
type Queue struct {
	//队列ID
	QueueId string

	//所属组ID
	GroupId string

	//队列容量
	Capacity int32

	//队列内消息类型
	QueueType int8

	//队列消息
	messages []message.Message

	//创建时间
	createTime time.Time

	//读写锁
	lock sync.RWMutex
}

// NewDelayQueue 构建一个默认的队列结构
func NewDelayQueue(queueID string, capacity int32) *Queue {
	queue := new(Queue)
	queue.QueueId = queueID
	queue.Capacity = capacity
	queue.createTime = time.Now()
	return queue
}

// Size 获取队列消息数量
func (q *Queue) Size() int32 {
	return int32(len(q.messages))
}

// Push 向队列添加消息 线程安全
func (q *Queue) Push(mes *message.Message) error {
	q.lock.RLock()
	defer q.lock.RUnlock()
	messages := q.messages
	//接收到的消息是否是队列的类型
	if getType(*mes) != q.QueueType {
		return errors.New("queue type inconsistent ")
	}
	if messages == nil {
		q.messages = make([]message.Message, 1, q.Capacity)
	}
	//如果队列不是无限容量，且有没足够容量时则抛出异常
	if q.Capacity != -1 && int32(len(q.messages)) >= q.Capacity {
		return errors.New("the queue capacity is full")
	}
	q.messages = append(q.messages, *mes)
	return nil
}

// Pull 拉取指定位置元素 线程安全
func (q *Queue) Pull(index int32) *message.Message {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.messages) == 0 {
		return nil
	}
	// 指定位置超过了元素的容量
	if index >= int32(len(q.messages)) {
		return nil
	}
	mess := q.messages[index]
	return &mess
}

// 判断消息接口的类型
func getType(mes any) int8 {
	switch mes.(type) {
	case message.RandomMessage:
		return enum.RandomMessage
	case message.DelayedMessage:
		return enum.DelayedMessage
	case message.ExclusiveMessage:
		return enum.ExclusiveMessage
	case message.WeightMessage:
		return enum.WeightMessage
	case message.SubMessage:
		return enum.SubMessage
	}
	return -1
}
