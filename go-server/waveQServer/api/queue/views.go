package queue

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waveQServer/comm"
	"waveQServer/core/groups"
	"waveQServer/core/queue/queueImpl"
	"waveQServer/identity"
	"waveQServer/utils/logutil"
)

func pull(c *gin.Context) {
	build := identity.BuilderUser().Build()
	messageType, message, err := c.Writer.(http.Hijacker).Hijack()
	err := c.ShouldBindJSON(build)
	if err != nil {
		logutil.LogError(err.Error())
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	que, err := groups.GetGroupQueueById(build.GroupId, build.QueueId)
	if err != nil {
		logutil.LogError(err.Error())
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	switch que.(type) {
	case *queueImpl.BroadcastQueue:
		message := que.(*queueImpl.BroadcastQueue).Pull(build.Index)
	}
}
