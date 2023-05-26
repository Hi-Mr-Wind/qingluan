package queue

import (
	"github.com/google/uuid"
	"time"
	"waveQServer/entity"
	"waveQServer/utils"
)

// Queue 队列结构
type Queue struct {
	//队列ID
	QueueId []byte

	//所属组ID
	GroupId []byte

	//队列容量
	Capacity int32

	//队列类型
	QueueType int8

	//队列模式 true为复制模式  否则为分发模式
	PatternCopy bool

	//队列消息
	Messages []entity.Message

	//创建时间
	CreateTime time.Time
}

// New 构建一个默认的队列结构
func New() *Queue {
	queue := new(Queue)
	queue.QueueId = []byte(uuid.New().String())
	queue.PatternCopy = false
	queue.Capacity = -1
	queue.QueueType = utils.STANDARD
	queue.CreateTime = time.Now()
	return queue
}

func NewQueue(queue Queue) *Queue {
	q := New()
	q.GroupId = queue.GroupId
	q.Capacity = queue.Capacity
	q.QueueType = queue.QueueType
	q.PatternCopy = queue.PatternCopy
	return q
}

// Push 向队列添加消息
func (q *Queue) Push(message *entity.Message) {
	messages := q.Messages
	if messages == nil {
		q.Messages = make([]entity.Message, 10, 20)
	}
	q.Messages = append(q.Messages, *message)
}

// Size 获取队列消息数量
func (q *Queue) Size() int {
	return len(q.Messages)
}

// Popup 弹出元素
func (q *Queue) Popup() (entity.Message, *[]entity.Message) {
	if len(q.Messages) == 0 {
		return entity.Message{}, &q.Messages
	}
	message := q.Messages[0]
	messages := q.Messages[1:]
	q.Messages = messages
	return message, &messages
}
