package repo

import (
	"context"
	"go-service/work/app/biz/domain/entity"
)

type UserRepo interface {
	FindByUid(ctx context.Context, uid string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)
	Save(ctx context.Context, account *entity.User) (*entity.User, error)
}
