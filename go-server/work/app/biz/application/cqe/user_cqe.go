package cqe

import (
	"go-service/pkg/errno"
	"go-service/work/app/biz/domain/vo"
)

type AccountUidQuery struct {
	Uid string `json:"uid"`
}

func (q *AccountUidQuery) Validate() error {
	if q.Uid == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "uid")
	}
	return nil
}

type AccountUsernameQuery struct {
	Username string `json:"uid"`
}

func (q *AccountUsernameQuery) Validate() error {
	if q.Username == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "username")
	}
	return nil
}

type CreateUserCommand struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (c *CreateUserCommand) Validate() error {
	if _, err := vo.NewUsername(c.Username).CheckFormat(); err != nil {
		return err
	}
	if _, err := vo.NewEmail(c.Email).CheckFormat(); err != nil {
		return err
	}
	if c.Phone != "" {
		if _, err := vo.NewPhoneNumber(c.Phone).CheckFormat(); err != nil {
			return err
		}
	}
	if c.Password != "" {
		if _, err := vo.NewPassword(c.Password).CheckFormat(); err != nil {
			return err
		}
	}
	return nil
}
