package dto

import "go-service/work/app/biz/domain/entity"

type UserDTO struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func UserToDTO(ent *entity.User) *UserDTO {
	return &UserDTO{
		Uid:      ent.Uid().GetUid(),
		Username: ent.Username().GetUsername(),
		Email:    ent.Email().GetEmail(),
		Phone:    ent.Phone().GetNumber(),
		Password: ent.Password().GetPlainPassword(),
	}
}
