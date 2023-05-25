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
	queue.PatternCopy = true
	queue.Capacity = -1
	queue.QueueType = utils.STANDARD
	queue.CreateTime = time.Now()
	return queue
}

// SetMessage 向队列添加消息
func (q *Queue) SetMessage(message *entity.Message) {
	messages := q.Messages
	if messages == nil {
		q.Messages = make([]entity.Message, 10, 20)
	}
	q.Messages = append(q.Messages, *message)
}
