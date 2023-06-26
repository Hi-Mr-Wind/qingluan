package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"waveQServer/src/comm"
	"waveQServer/src/comm/req/cqe"
	"waveQServer/src/core/cache"
	"waveQServer/src/core/database"
	"waveQServer/src/core/groups"
	"waveQServer/src/entity"
	"waveQServer/src/utils"
	"waveQServer/src/utils/logutil"
)

// Login 管理员登录
func Login(c *gin.Context) {
	admin := &cqe.AdminCmd{}
	if err := c.ShouldBindJSON(admin); err != nil {
		logutil.LogError(err.Error())
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	if err := admin.Validate(); err != nil {
		logutil.LogError(err.Error())
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		return
	}
	data, code := admin.Login()
	c.JSON(code, data)
	return
}

// CreateApiKey 创建消费者apikey
func CreateApiKey(c *gin.Context) {
	r := &cqe.CreateApiKeyCmd{}
	err := c.ShouldBindJSON(r)
	if err != nil {
		comm.DisposeError(err, c)
		return
	}
	if err := r.Validate(); err != nil {
		logutil.LogError(err.Error())
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		return
	}
	user := new(entity.User)
	user.ApiKey = utils.GetApiKey(r.RecessRights)
	user.Id = uuid.New().String()
	user.CreatTime = utils.GetTime()
	user.ExpirationTime = r.ExpirationTime
	us := make([]entity.QueueUser, 1, 10)
	for _, v := range r.RecessRights {
		u := new(entity.QueueUser)
		u.Id = uuid.New().String()
		u.QueueId = v
		u.UserId = user.Id
		us = append(us, *u)
	}
	cache.AddApikey(user, r.RecessRights)
	//异步持久化
	go func() {
		database.GetDb().Create(user)
		database.GetDb().Create(us)
	}()
	c.JSON(http.StatusOK, comm.OK(user))
	return
}

// CreateGroup 创建一个组
func CreateGroup(c *gin.Context) {
	group := make(map[string]string)
	err := c.ShouldBindJSON(group)
	if err != nil {
		comm.DisposeError(err, c)
		return
	}
	if group["groupId"] == "" {
		comm.DisposeError(errors.New("groupId is null"), c)
		return
	}
	if groups.GetGroupById(group["groupId"]) != nil {
		comm.DisposeError(errors.New("the group is exist"), c)
		return
	}
	_, err = groups.NewGroup(group["groupId"])
	if err != nil {
		comm.DisposeError(err, c)
		return
	}
	c.JSON(http.StatusOK, comm.OK())
	return
}
