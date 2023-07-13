package app

import (
	"context"
	"go-service/pkg/assert"
	"go-service/work/app/biz/application/cqe"
	"go-service/work/app/biz/application/dto"
	"go-service/work/app/biz/domain/entity"
	"go-service/work/app/biz/domain/repo"
	"go-service/work/app/biz/domain/service"
	"go-service/work/app/biz/infrastructure/dao"
	"go-service/work/app/biz/infrastructure/persistence"
	"go-service/work/internal/resource"
	"sync"
)

var (
	userAppOnce      sync.Once
	singletonUserApp UserApp
)

type UserApp interface {
	GetByUid(ctx context.Context, query *cqe.AccountUidQuery) (*dto.UserDTO, error)
	GetByUsername(ctx context.Context, query *cqe.AccountUsernameQuery) (*dto.UserDTO, error)
	CreateUser(ctx context.Context, cmd *cqe.CreateUserCommand) (*dto.UserDTO, error)
}

type userApp struct {
	userRepo repo.UserRepo
	userSrv  service.UserService
}

// DefaultUserApp 默认单例构建方法
func DefaultUserApp() UserApp {
	assert.NotCircular()
	userAppOnce.Do(func() {
		var (
			db       = resource.DefaultMysqlResource().RwRepo()
			userDao  = dao.NewUserDao(db)
			userRepo = persistence.NewUserRepo(userDao)
		)
		singletonUserApp = &userApp{
			userRepo: userRepo,
			//bizChannelRepo: bizChannelRepo,
			//userSrv:        service.NewUserService(userRepo),
		}
	})
	assert.NotNil(singletonUserApp)
	return singletonUserApp
}

func (u *userApp) GetByUid(ctx context.Context, query *cqe.AccountUidQuery) (*dto.UserDTO, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}
	// TODO 待完成
	return nil, nil
}

func (u *userApp) GetByUsername(ctx context.Context, query *cqe.AccountUsernameQuery) (*dto.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userApp) CreateUser(ctx context.Context, cmd *cqe.CreateUserCommand) (*dto.UserDTO, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	account := entity.NewUserWithDefault(cmd.Username, cmd.Password, cmd.Email, cmd.Phone)

	// 检查账号信息是否已经存在
	err := u.userSrv.CheckUserExists(ctx, account)
	if err != nil {
		return nil, err
	}

	// 创建账号
	m, err := u.userRepo.Save(ctx, account)
	if err != nil {
		return nil, err
	}

	return dto.UserToDTO(m), nil
}
