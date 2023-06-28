package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"waveQServer/src/comm"
	"waveQServer/src/comm/req/cqe"
	"waveQServer/src/core/cache"
	"waveQServer/src/core/groups"
	"waveQServer/src/core/message"
	"waveQServer/src/utils"
	"waveQServer/src/utils/httpUtils"
	"waveQServer/src/utils/logutil"
)

// Pull 拉取消息
func Pull(c *gin.Context) {
	query := &cqe.QueueInfoQuery{
		QueueId: c.Query("queue_id"),
		GroupId: c.Query("group_id"),
	}
	if err := query.Validate(); err != nil {
		logutil.LogError(err.Error())
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		return
	}
	apiKey := cqe.GetApiKey(c)
	//判断用户是否有足够权限
	if cache.IsValidLimit(apiKey, query.QueueId) {
		comm.DisposeError(errors.New("insufficient privileges"), c)
		return
	}
	queue, err := groups.GetGroupQueueById(query.GroupId, query.QueueId)
	if err != nil {
		logutil.LogInfo(err.Error())
		comm.DisposeError(err, c)
		return
	}
	//每100毫秒查询一次消息，如果查询100次后（10秒）依旧没有消息则返回null
	for i := 0; i < 100; i++ {
		mes := queue.Pull(utils.JsonToMap(cache.GetUser(apiKey).Answer)[query.QueueId].(int32))
		if mes != nil {
			c.JSON(http.StatusOK, comm.OK(mes))
			c.Abort()
			//跳转到退出位置，防止变量污染
			goto exit
		}
		<-time.After(100 * time.Millisecond)
		continue
	}
	//超时未能拿到消息，则返回空
	c.JSON(http.StatusOK, comm.OK(nil))
exit:
	return
}

// Push 添加消息
func Push(c *gin.Context) {
	queueMessage := &cqe.QueueMessageInfoCmd{}
	if err := c.ShouldBindJSON(queueMessage); err != nil {
		logutil.LogError(err.Error())
		comm.DisposeError(err, c)
		return
	}
	if err := queueMessage.Validate(); err != nil {
		logutil.LogError(err.Error())
		comm.DisposeError(err, c)
		return
	}
	// 获取队列信息
	queue, err := groups.GetGroupQueueById(queueMessage.GroupId, queueMessage.QueueId)
	if err != nil {
		logutil.LogError(err.Error())
		comm.DisposeError(err, c)
		return
	}
	// 验证token
	httpUtils.Token(c)
	newMessage := message.NewMessage(queueMessage)
	if newMessage == nil {
		logutil.LogError("消息类型不符合规定")
		comm.DisposeError(errors.New("the message type is invalid"), c)
		return
	}
	err = queue.Push(&newMessage)
	if err != nil {
		logutil.LogInfo(err.Error())
		comm.DisposeError(err, c)
		return
	}
	c.JSON(http.StatusOK, comm.OK())
}
