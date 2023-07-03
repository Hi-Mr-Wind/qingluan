package cqe

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type QueueInfoQuery struct {
	QueueId string `json:"queue_id"`
	GroupId string `json:"group_id"`
}

func (q *QueueInfoQuery) Validate() error {
	if len(q.QueueId) == 0 {
		return fmt.Errorf("queue_id is empty")
	}
	if len(q.GroupId) == 0 {
		return fmt.Errorf("group_id is empty")
	}
	return nil
}

type QueueMessageInfoCmd struct {
	QueueId     string   `json:"queue_id"`
	GroupId     string   `json:"group_id"`
	ProducerId  string   `json:"producer_id"`
	MessageType int8     `json:"message_type"`
	Indate      int64    `json:"indate"`
	Body        []byte   `json:"body"`
	Subscriber  []string `json:"subscriber"`       // 订阅消息字段
	Weight      int32    `json:"weight,omitempty"` // 随机消息or权重消息字段：权重消息
	Number      int32    `json:"number,omitempty"` // 随机消息字段：可消费次数
	Delayed     int64    `json:"delayed"`          // 延迟消息字段
}

func (q *QueueMessageInfoCmd) Validate() error {
	if len(q.QueueId) == 0 {
		return fmt.Errorf("queue_id is empty")
	}
	if len(q.GroupId) == 0 {
		return fmt.Errorf("group_id is empty")
	}
	if q.Body == nil {
		return fmt.Errorf("body is empty")
	}
	return nil
}

func GetApiKey(c *gin.Context) string {
	return c.Request.Header.Get("API_KEY")
}

func GetToken(c *gin.Context) string {
	return c.Request.Header.Get("token")
}
