package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"waveQServer/src/comm"
	"waveQServer/src/core/cache"
	"waveQServer/src/core/groups"
	"waveQServer/src/utils"
	"waveQServer/src/utils/logutil"
)

// Pull 拉取消息
func Pull(c *gin.Context) {
	queueId := c.Query("queueId")
	groupId := c.Query("groupId")
	apiKey := c.Request.Header.Get("API_KEY")
	//判断用户是否有足够权限
	if cache.IsValidLimit(apiKey, queueId) {
		comm.DisposeError(errors.New("insufficient privileges"), c)
		return
	}
	queue, err := groups.GetGroupQueueById(groupId, queueId)
	if err != nil {
		logutil.LogInfo(err.Error())
		comm.DisposeError(err, c)
		return
	}
	//每100毫秒查询一次消息，如果查询100次后（10秒）依旧没有消息则返回null
	for i := 0; i < 100; i++ {
		message := queue.Pull(utils.JsonToMap(cache.GetUser(apiKey).Answer)[queueId].(int32))
		if message != nil {
			c.JSON(http.StatusOK, comm.OK(message))
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
