package persistence

import (
	"context"
	"go-service/work/app/biz/domain/entity"
	"go-service/work/app/biz/infrastructure/convertor"
	"go-service/work/app/biz/infrastructure/dao"
)

type userRepoImpl struct {
	userDao *dao.UserDao
}

func NewUserRepo(userDao *dao.UserDao) *userRepoImpl {
	return &userRepoImpl{
		userDao: userDao,
	}
}

func (r *userRepoImpl) FindByUid(ctx context.Context, uid string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepoImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepoImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepoImpl) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepoImpl) Save(ctx context.Context, user *entity.User) (*entity.User, error) {
	userInfo := convertor.ToUserPo(user)

	if userInfo.Id > 0 {
		_, err := r.userDao.Update(ctx, userInfo)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := r.userDao.Insert(ctx, userInfo)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
