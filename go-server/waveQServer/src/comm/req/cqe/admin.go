package cqe

import (
	"fmt"
	"net/http"
	"waveQServer/src/comm"
	"waveQServer/src/core/database"
	"waveQServer/src/entity"
	"waveQServer/src/utils"
	"waveQServer/src/utils/jwtutil"
	"waveQServer/src/utils/logutil"
)

// AdminCmd admin用户登录参数
type AdminCmd struct {
	UserName string `json:"column:user_name"`
	Password string `json:"column:password"`
}

// Validate 校验登录参数
func (a *AdminCmd) Validate() error {
	if utils.IsEmpty(a.UserName) {
		return fmt.Errorf("UserName is empty")
	}
	if utils.IsEmpty(a.Password) {
		return fmt.Errorf("PassWord is empty")
	}
	return nil
}

// Login 管理员登录校验
func (a *AdminCmd) Login() (*comm.JsonResult, int) {
	adm := new(entity.Admin)
	md5 := utils.Md5([]byte(a.UserName))
	database.GetDb().Where("user_name = ?", md5).Find(adm)
	if utils.IsEmpty(adm.Id) {
		logutil.LogInfo("login failure! The account or password is incorrect.")
		return comm.Fail("login failure! The account or password is incorrect."), http.StatusBadRequest
	}
	if utils.NotEquals(adm.UserName, utils.Md5([]byte(a.UserName))) {
		return comm.Fail("username error"), http.StatusBadRequest
	}
	if utils.NotEquals(adm.Password, utils.Md5([]byte(a.Password))) {
		return comm.Fail("password error"), http.StatusBadRequest
	}
	token, err := jwtutil.GetToken(a.UserName, a.Password)
	if err != nil {
		return comm.Fail(err.Error()), http.StatusBadRequest
	}
	m := make(map[string]string)
	m["XMD-TOKEN"] = token
	return comm.OK(m), http.StatusOK
}

// CreateApiKeyCmd 创建apikey接收参数
type CreateApiKeyCmd struct {
	//权限
	RecessRights []string
	//过期时间
	ExpirationTime int64
}

func (c *CreateApiKeyCmd) Validate() error {
	if len(c.RecessRights) == 0 {
		return fmt.Errorf("RecessRights is empty")
	}
	return nil
}

// 创建token
type CreateTokenCmd struct {
	UserId     string   `json:"user_id"`    // 消费者ID
	Permission []string `json:"permission"` // token的权限
}

func (c *CreateTokenCmd) Validate() error {
	if len(c.UserId) == 0 {
		return fmt.Errorf("user_id is empty")
	}
	return nil
}

// 删除token
type DeleteTokenCmd struct {
	Token string `json:"token"`
}

func (c *DeleteTokenCmd) Validate() error {
	if len(c.Token) == 0 {
		return fmt.Errorf("token is empty")
	}
	return nil
}
