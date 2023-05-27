package user

import (
	"errors"
	"time"
	"waveQServer/config"
	"waveQServer/utils"
	"waveQServer/utils/logutil"
)

// User 临时用户结构
type User struct {
	//apikey
	ApiKey string `json:"apiKey"`
	//访问队列权限
	RccessRights [][]byte `json:"rccessRights"`
	//过期时间
	OutTime time.Time
}

// Admin 管理员用户对象
type Admin struct {
	//用户名
	UserName string `json:"userName"`
	//密码
	Password string `json:"password"`
}

// NewUser 创建一个新的临时用户对象
func NewUser() *User {
	return new(User)
}

// NewAdmin 创建一个新的管理员用户对象
func NewAdmin() *Admin {
	admin := new(Admin)
	admin.UserName = config.GetConfig().UserName
	admin.Password = config.GetConfig().Password
	return admin
}

// setApiKey 设置一个APIkey
func (u *User) setApiKey() error {
	if len(u.RccessRights) == 0 {
		s := "you cannot create an unauthorized user"
		logutil.LogWarning(s)
		return errors.New(s)
	}
	u.ApiKey = utils.GetApiKey(u.RccessRights)
	return nil
}

// setOutTime 设置用户过期时间
func (u *User) setOutTime(date string) {
	parse, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		logutil.LogError(err.Error())
		return
	}
	u.OutTime = parse
}
