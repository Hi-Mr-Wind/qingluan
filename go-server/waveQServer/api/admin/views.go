package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waveQServer/comm"
	"waveQServer/config"
	"waveQServer/identity"
	"waveQServer/utils"
	"waveQServer/utils/jwtutil"
	"waveQServer/utils/logutil"
)

// Login 管理员登录
func Login(c *gin.Context) {
	admin := identity.NewAdmin()
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		logutil.LogError(err.Error())
		fail := comm.Fail(err.Error())
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	if utils.NotEquals(config.GetConfig().UserName, utils.Md5([]byte(admin.UserName))) {
		fail := comm.Fail("username error")
		c.JSON(http.StatusBadRequest, fail)
		c.Abort()
		return
	}
	if utils.NotEquals(config.GetConfig().Password, utils.Md5([]byte(admin.Password))) {
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
