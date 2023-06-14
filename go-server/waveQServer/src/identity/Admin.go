package identity

import (
	"waveQServer/src/config"
)

// Admin 管理员用户对象
type Admin struct {
	//用户名
	UserName string `json:"userName"`
	//密码
	Password string `json:"password"`
}

// NewAdmin 创建一个新的管理员用户对象
func NewAdmin() *Admin {
	admin := new(Admin)
	admin.UserName = config.GetConfig().UserName
	admin.Password = config.GetConfig().Password
	return admin
}
