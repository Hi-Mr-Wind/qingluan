package gropu

import (
	"errors"
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

// NewGroup 构造一个组对象
func NewGroup(groupId []byte) (*Group, error) {
	if _, ok := groups[string(groupId)]; ok {
		return nil, errors.New("the groupId is already existed")
	}
	group := new(Group)
	group.GroupId = groupId
	group.GroupQueue = make(map[string]*queue.Queue, 50)
	groups[string(group.GroupId)] = group
	return group, nil
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
func (g *Group) BindQueue(que *queue.Queue) error {
	q := g.GroupQueue[string(que.QueueId)]
	if q != nil {
		return errors.New("the queue is already in the group")
	}
	g.GroupQueue[string(que.QueueId)] = que
	return nil
}