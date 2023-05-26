package gropu

import (
	"errors"
	"github.com/google/uuid"
	"waveQServer/entity/queue"
)

// Groups 组对象集
var groups = make(map[string]*Group, 50)

// Group 组结构
type Group struct {

	// 组ID
	GroupId []byte
	// 组内队列
	GroupQueue map[string]*queue.Queue
}

// New 构造一个组对象
func New() *Group {
	group := new(Group)
	group.GroupId = []byte(uuid.New().String())
	group.GroupQueue = make(map[string]*queue.Queue, 50)
	groups[string(group.GroupId)] = group
	return group
}

// GetGroupById 根据组ID获取一个组对象
func GetGroupById(id []byte) *Group {
	return groups[string(id)]
}

// GetGroupQueueById 根据队列ID获取队列
func (g *Group) GetGroupQueueById(queueId []byte) (*queue.Queue, error) {
	q, ok := g.GroupQueue[string(queueId)]
	if !ok {
		return nil, errors.New("the queue is not in Group")
	}
	return q, nil
}

// BindQueue 向组中添加一个队列
func (g *Group) BindQueue(que *queue.Queue) {
	g.GroupQueue[string(que.QueueId)] = que
}
