package cqe

import "fmt"

type AdminCmd struct {
	UserName string `json:"column:user_name"`
	Password string `json:"column:password"`
}

func (a *AdminCmd) Validate() error {
	if a.UserName == "" {
		return fmt.Errorf("UserName is empty")
	}
	if a.Password == "" {
		return fmt.Errorf("PassWord is empty")
	}
	return nil
}

// CreateApiKeyReq 创建apikey接收参数
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
