package convertor

import (
	"go-service/work/app/biz/domain/entity"
	"go-service/work/app/biz/infrastructure/dao/po"
)

func ToUserEntity(user *po.UserInfoModel) *entity.User {
	return entity.NewUser(
		user.Id,
		user.Uid,
		user.UserName,
		user.Email,
		user.Phone,
		user.Salt,
		user.Credential,
		user.RegisterTime,
	)
}

func ToUserPo(user *entity.User) *po.UserInfoModel {
	userInfoPo := &po.UserInfoModel{
		BaseModel:    po.BaseModel{Id: user.Id().GetId()},
		Uid:          user.Uid().GetUid(),
		UserName:     user.Username().GetUsername(),
		Email:        user.Email().GetEmail(),
		Phone:        user.Phone().GetNumber(),
		Salt:         user.Password().GetSalt(),
		RegisterTime: user.RegisterTime().GetTime(),
	}
	return userInfoPo
}
