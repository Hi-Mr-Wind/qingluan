package queueImpl

import (
	"errors"
	"sync"
	"time"
	"waveQServer/entity/message"
	"waveQServer/identity"
	"waveQServer/utils/lastingUtils"
)

// StandardQueue 标准队列结构
type StandardQueue struct {
	//队列ID
	QueueId string

	//所属组ID
	GroupId string

	//队列容量
	Capacity int32

	//队列消息
	messages []message.Message

	//消费者
	monitor map[string]*identity.User

	//创建时间
	createTime time.Time

	//定义写锁
	lock sync.RWMutex
}

// NewStandardQueue 构建一个默认的队列结构
func NewStandardQueue(queueID []byte) (*StandardQueue, error) {
	queue := new(StandardQueue)
	queue.QueueId = string(queueID)
	queue.Capacity = 10000
	queue.createTime = time.Now()
	return queue, nil
}

// Size 获取队列消息数量
func (q *StandardQueue) Size() int32 {
	return int32(len(q.messages))
}

// Push 向队列添加消息 线程安全
func (q *StandardQueue) Push(mes *message.Message) {
	q.lock.Lock()
	defer q.lock.Unlock()
	messages := q.messages
	if messages == nil {
		q.messages = make([]message.Message, 1, q.Capacity)
	}
	if int32(len(q.messages)) >= q.Capacity {

	}
	//获取前条消息的ID
	e := q.messages[len(q.messages)-1]
	mes.Header.FormerId = e.Header.Id
	// 异步将消息持久化
	go lastingUtils.AsyncMessage(mes)
	q.messages = append(q.messages, *mes)
}

// Pull 拉取并删除最先进入的元素 线程安全
func (q *StandardQueue) Pull() (*message.Message, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.messages) == 0 {
		return nil, errors.New("the queue is empty")
	}
	message := q.messages[0]
	q.messages = q.messages[1:]
	return &message, nil
}

// AddUser 添加一个队列消费者
func (q *StandardQueue) AddUser(user *identity.User) {
	q.monitor[user.ApiKey] = user
}

func (q *StandardQueue) GetQueueId() string {
	return q.QueueId
}
