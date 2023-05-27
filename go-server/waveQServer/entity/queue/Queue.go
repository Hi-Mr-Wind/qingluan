package queue

import (
	"errors"
	"sync"
	"time"
	"waveQServer/entity"
	"waveQServer/utils"
)

// Groups 组对象集
var groups = make(map[string]*Group, 50)

// Group 组结构
type Group struct {

	// 组ID
	GroupId []byte
	// 组内队列
	GroupQueue map[string]*Queue
}

// Queue 队列结构
type Queue struct {
	//队列ID
	QueueId []byte

	//所属组ID
	GroupId []byte

	//队列容量
	capacity int32

	//队列类型
	queueType int8

	//队列模式 true为复制模式  否则为分发模式
	patternCopy bool

	//队列消息
	messages []entity.Message

	//创建时间
	createTime time.Time

	//定义写锁
	lock sync.RWMutex
}

// NewGroup 构造一个组对象
func NewGroup(groupId []byte) (*Group, error) {
	if _, ok := groups[string(groupId)]; ok {
		return nil, errors.New("the groupId is already existed")
	}
	group := new(Group)
	group.GroupId = groupId
	group.GroupQueue = make(map[string]*Queue, 50)
	groups[string(group.GroupId)] = group
	return group, nil
}

// New 构建一个默认的队列结构,参数1所属群组，参数2 队列ID
func New(groupId []byte, queueID []byte) (*Queue, error) {
	gro := GetGroupById(groupId)
	if gro == nil {
		return nil, errors.New("this group is undefined")
	}
	queue := new(Queue)
	queue.QueueId = queueID
	queue.GroupId = groupId
	queue.patternCopy = false
	queue.capacity = 10000
	queue.queueType = utils.STANDARD
	queue.createTime = time.Now()
	//向组内添加队列
	err := gro.BindQueue(queue)
	if err != nil {
		return nil, err
	}
	return queue, nil
}

// SetCapacity 设置队列容量
func (q *Queue) SetCapacity(capacity int32) {
	q.capacity = capacity
}

// SetQueueType 设置队列类型
func (q *Queue) SetQueueType(queueType int8) {
	q.queueType = queueType
}

// SetPatternCopy 设置队列模式 true为复制模式  否则为分发模式
func (q *Queue) SetPatternCopy(patternCopy bool) {
	q.patternCopy = patternCopy
}

// Size 获取队列消息数量
func (q *Queue) Size() int32 {
	return int32(len(q.messages))
}

// Push 向队列添加消息 线程安全
func (q *Queue) Push(message *entity.Message) {
	q.lock.Lock()
	defer q.lock.Unlock()
	messages := q.messages
	if messages == nil {
		q.messages = make([]entity.Message, 1, q.capacity)
	}
	if int32(len(q.messages)) >= q.capacity {

	}
	//获取前条消息的ID
	e := q.messages[len(q.messages)-1]
	message.Header.FormerId = e.Header.Id
	q.messages = append(q.messages, *message)
}

// Pull 拉取并删除最先进入的元素 线程安全
func (q *Queue) Pull() (*entity.Message, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.messages) == 0 {
		return nil, errors.New("the queue is empty")
	}
	message := q.messages[0]
	q.messages = q.messages[1:]
	return &message, nil
}

// PullByIndex 获取队列中指定下标的元素
func (q *Queue) PullByIndex(index int32) (*entity.Message, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if index > int32(len(q.messages)-1) {
		return nil, errors.New("array index out of bounds")
	}
	return &q.messages[index], nil
}

// GetGroupById 根据组ID获取一个组对象
func GetGroupById(id []byte) *Group {
	return groups[string(id)]
}

// GetGroupQueueById 根据队列ID获取队列
func (g *Group) GetGroupQueueById(queueId []byte) (*Queue, error) {
	q, ok := g.GroupQueue[string(queueId)]
	if !ok {
		return nil, errors.New("the queue is not in Group")
	}
	return q, nil
}

// BindQueue 向组中添加一个队列
func (g *Group) BindQueue(que *Queue) error {
	q := g.GroupQueue[string(que.QueueId)]
	if q != nil {
		return errors.New("the queue is already in the group")
	}
	g.GroupQueue[string(que.QueueId)] = que
	return nil
}
