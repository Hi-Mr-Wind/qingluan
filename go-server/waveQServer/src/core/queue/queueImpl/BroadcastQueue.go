package queueImpl

import (
	"sync"
	"time"
	"waveQServer/src/core/message"
	"waveQServer/src/entity"
	"waveQServer/src/utils/logutil"
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
	messages []message.Message

	//消费者
	monitor map[string]*entity.User

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
func (q *BroadcastQueue) Push(mes *message.Message) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	messages := q.messages
	if messages == nil {
		q.messages = make([]message.Message, 1, q.Capacity)
	}
	if int32(len(q.messages)) >= q.Capacity {

	}
	// 异步将消息持久化
	go message.SetCachedSubMessage(mes)
	q.messages = append(q.messages, *mes)
}

// Pull 拉取元素 线程安全
func (q *BroadcastQueue) Pull(index int32) *message.Message {
	q.lock.Lock()
	defer q.lock.Unlock()
	if index >= q.Size() {
		logutil.LogInfo(q.GroupId + " --queue:" + q.QueueId + "is nonentity message ")
		return nil
	}
	entity.User{}.MessageChan <- &q.messages[index]
	return &q.messages[index]
}

// AddUser 添加一个队列消费者
func (q *BroadcastQueue) AddUser(user *entity.User) {
	q.monitor[user.ApiKey] = user
}

func (q *BroadcastQueue) GetQueueId() string {
	return q.QueueId
}
