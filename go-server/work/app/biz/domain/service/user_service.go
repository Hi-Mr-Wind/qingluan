package service

import (
	"context"
	"go-service/pkg/errno"
	"go-service/work/app/biz/domain/entity"
	"go-service/work/app/biz/domain/repo"
)

type UserService interface {
	CheckUserExists(ctx context.Context, user *entity.User) error
}

type userServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) *userServiceImpl {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) CheckUserExists(ctx context.Context, user *entity.User) error {
	existed, err := s.userRepo.FindByUsername(ctx, user.Username().GetUsername())
	if err != nil {
		return err
	}
	if existed != nil {
		if user.Id().GetId() == 0 || user.Id().GetId() != existed.Id().GetId() {
			return errno.NewSimpleBizError(errno.ErrUsernameExists, nil)
		}
	}

	if user.Email().GetEmail() != "" {
		existed, err = s.userRepo.FindByEmail(ctx, user.Email().GetEmail())
		if err != nil {
			return err
		}
		if existed != nil {
			if user.Id().GetId() == 0 || user.Id().GetId() != existed.Id().GetId() {
				return errno.NewSimpleBizError(errno.ErrEmailExists, nil)
			}
		}
	}

	if user.Phone().GetNumber() != "" {
		existed, err = s.userRepo.FindByPhone(ctx, user.Phone().GetNumber())
		if err != nil {
			return err
		}
		if existed != nil {
			if user.Id().GetId() == 0 || user.Id().GetId() != existed.Id().GetId() {
				return errno.NewSimpleBizError(errno.ErrPhoneExists, nil)
			}
		}
	}

	return nil
}
