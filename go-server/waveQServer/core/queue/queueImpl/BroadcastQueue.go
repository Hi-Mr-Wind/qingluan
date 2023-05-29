package queueImpl

import (
	"sync"
	"time"
	"waveQServer/entity"
	"waveQServer/identity"
	"waveQServer/utils"
)

// BroadcastQueue 广播队列结构
type BroadcastQueue struct {
	//队列ID
	QueueId string

	//所属组ID
	GroupId string

	//队列容量
	Capacity int32

	//队列消息
	messages []entity.Message

	//消费者
	monitor []*identity.User

	//创建时间
	createTime time.Time

	//定义写锁
	lock sync.RWMutex
}

// NewBroadcastQueue 构建一个默认的队列结构
func NewBroadcastQueue(queueID []byte) (*BroadcastQueue, error) {
	queue := new(BroadcastQueue)
	queue.QueueId = string(queueID)
	queue.Capacity = 10000
	queue.createTime = time.Now()
	return queue, nil
}

// Size 获取队列消息数量
func (q *BroadcastQueue) Size() int32 {
	return int32(len(q.messages))
}

// Push 向队列添加消息 线程安全
func (q *BroadcastQueue) Push(message *entity.Message) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	messages := q.messages
	if messages == nil {
		q.messages = make([]entity.Message, 1, q.Capacity)
	}
	if int32(len(q.messages)) >= q.Capacity {

	}
	//获取前条消息的ID
	e := q.messages[len(q.messages)-1]
	message.Header.FormerId = e.Header.Id
	// 异步将消息持久化
	go utils.AsyncMessage(message)
	q.messages = append(q.messages, *message)
}

// Pull 拉取元素 线程安全
func (q *BroadcastQueue) Pull(index int32) *entity.Message {
	q.lock.Lock()
	defer q.lock.Unlock()
	if index >= q.Size() {
		return nil
	}
	return &q.messages[index]
}

// AddUser 添加一个队列消费者
func (q *BroadcastQueue) AddUser(user *identity.User) {
	q.monitor = append(q.monitor, user)
}
