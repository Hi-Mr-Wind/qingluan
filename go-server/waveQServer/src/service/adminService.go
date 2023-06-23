package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waveQServer/src/comm"
	"waveQServer/src/core/database"
	"waveQServer/src/core/groups"
	"waveQServer/src/entity"
	"waveQServer/src/utils"
	"waveQServer/src/utils/jwtutil"
	"waveQServer/src/utils/logutil"
)

// Login 管理员登录
func Login(c *gin.Context) {
	admin := entity.NewAdmin()
	err := c.ShouldBindJSON(admin)
	if err != nil {
		logutil.LogError(err.Error())
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	adm := new(entity.Admin)
	md5 := utils.Md5([]byte(admin.UserName))
	database.GetDb().Where("user_name = ?", md5).Find(&adm)
	if utils.IsEmpty(adm.Id) {
		logutil.LogInfo("login failure! The account or password is incorrect.")
		fail := comm.Fail("login failure! The account or password is incorrect.")
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	if utils.NotEquals(adm.UserName, utils.Md5([]byte(admin.UserName))) {
		fail := comm.Fail("username error")
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	if utils.NotEquals(adm.Password, utils.Md5([]byte(admin.Password))) {
		fail := comm.Fail("password error")
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	token, err := jwtutil.GetToken(admin.UserName, admin.Password)
	if err != nil {
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	m := make(map[string]string)
	m["XMD-TOKEN"] = token
	ok := comm.OK(m)
	c.JSON(http.StatusOK, ok)
	c.Abort()
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
	_, err = groups.NewGroup(group["groupId"])
	if err != nil {
		comm.DisposeError(err, c)
		return
	}
	c.JSON(http.StatusOK, comm.OK())
	return
}
