package entity

import (
	"go-service/work/app/biz/domain/vo"
	"time"
)

type User struct {
	id           vo.ID
	uid          vo.UserUid
	username     vo.Username
	email        vo.Email
	phone        vo.PhoneNumber
	password     vo.Password
	registerTime vo.Time
}

func (a *User) Id() vo.ID {
	return a.id
}

func (a *User) Uid() vo.UserUid {
	return a.uid
}

func (a *User) Username() vo.Username {
	return a.username
}

func (a *User) Email() vo.Email {
	return a.email
}

func (a *User) Phone() vo.PhoneNumber {
	return a.phone
}

func (a *User) Password() vo.Password {
	return a.password
}

func (a *User) RegisterTime() vo.Time {
	return a.registerTime
}

func NewUser(id int64, uid, username, email, phone, salt, credential string, registerTime time.Time) *User {
	return &User{
		id:           vo.NewID(id),
		uid:          vo.NewUserUid(uid),
		username:     vo.NewUsername(username),
		email:        vo.NewEmail(email),
		phone:        vo.NewPhoneNumber(phone),
		password:     vo.NewEncryptedPassword(salt, credential),
		registerTime: vo.NewTime(registerTime),
	}
}

func NewUserWithDefault(username, password, email, phone string) *User {
	// 允许初始密码为空，若为空则自动生成随机密码
	var pwd vo.Password
	if password == "" {
		pwd = vo.NextRandomPassword().Encrypt()
	} else {
		pwd = vo.NewPassword(password).Encrypt()
	}

	return &User{
		uid:          vo.NextUserUid(),
		username:     vo.NewUsername(username),
		email:        vo.NewEmail(email),
		phone:        vo.NewPhoneNumber(phone),
		password:     pwd,
		registerTime: vo.NewTime(time.Now()),
	}
}
