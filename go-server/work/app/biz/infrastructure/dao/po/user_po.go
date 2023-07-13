package po

import "time"

type UserInfoModel struct {
	BaseModel
	Uid          string
	UserName     string
	Email        string
	Phone        string
	Salt         string
	Credential   string
	RegisterTime time.Time
}

func (m *UserInfoModel) TableName() string {
	return "user_info"
}
