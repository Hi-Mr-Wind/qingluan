package queueImpl

import (
	"errors"
	"sync"
	"time"
	"waveQServer/src/core/message"
	"waveQServer/src/entity"
)

// DelayQueue 延迟队列结构
type DelayQueue struct {
	//队列ID
	QueueId string

	//所属组ID
	GroupId string

	//队列容量
	Capacity int32

	//队列消息
	messages []message.Message

	//消费者
	monitor map[string]*entity.User

	//创建时间
	createTime time.Time

	//定义写锁
	lock sync.RWMutex
}

// NewDelayQueue 构建一个默认的队列结构
func NewDelayQueue(queueID []byte) (*DelayQueue, error) {
	queue := new(DelayQueue)
	queue.QueueId = string(queueID)
	queue.Capacity = 10000
	queue.createTime = time.Now()
	return queue, nil
}

// Size 获取队列消息数量
func (q *DelayQueue) Size() int32 {
	return int32(len(q.messages))
}

// Push 向队列添加消息 线程安全
func (q *DelayQueue) Push(mes *message.Message) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	messages := q.messages
	if messages == nil {
		q.messages = make([]message.Message, 1, q.Capacity)
	}
	if int32(len(q.messages)) >= q.Capacity {

	}
	q.messages = append(q.messages, *mes)
}

// Pull 拉取并删除最先进入的元素 线程安全
func (q *DelayQueue) Pull() (*message.Message, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.messages) == 0 {
		return nil, errors.New("the queue is empty")
	}
	mess := q.messages[0]
	q.messages = q.messages[1:]
	return &mess, nil
}

// AddUser 添加一个队列消费者
func (q *DelayQueue) AddUser(user *entity.User) {
	q.monitor[user.ApiKey] = user
}

func (q *DelayQueue) GetQueueId() string {
	return q.QueueId
}
