package dao

import (
	"context"
	"errors"
	"go-service/pkg/repository"
	"go-service/work/app/biz/infrastructure/dao/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDao struct {
	*baseDao
}

func NewUserDao(repo *repository.Database) *UserDao {
	return &UserDao{newBaseDao(repo)}
}

func (d *UserDao) SelectByUid(ctx context.Context, uid string) (*po.UserInfoModel, error) {
	var m po.UserInfoModel
	err := d.db.Self.Where("uid = ? AND user_status != 'deleted'", uid).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		d.logger.Errorf("SelectByUid error: %v", err)
		return nil, d.bizError(err)
	}
	return &m, nil
}

func (d *UserDao) SelectByUsername(ctx context.Context, username string) (*po.UserInfoModel, error) {
	var m po.UserInfoModel
	err := d.db.Self.Where("user_name = ? AND user_status != 'deleted'", username).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		d.logger.Errorf("SelectByUsername error: %v", err)
		return nil, d.bizError(err)
	}
	return &m, nil
}

func (d *UserDao) SelectByEmail(ctx context.Context, email string) (*po.UserInfoModel, error) {
	var m po.UserInfoModel
	err := d.db.Self.Where("email = ? AND user_status != 'deleted'", email).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		d.logger.Errorf("SelectByEmail error: %v", err)
		return nil, d.bizError(err)
	}
	return &m, nil
}

func (d *UserDao) SelectByPhone(ctx context.Context, phone string) (*po.UserInfoModel, error) {
	var m po.UserInfoModel
	err := d.db.Self.Where("phone = ? AND user_status != 'deleted'", phone).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		d.logger.Errorf("SelectByPhone error: %v", err)
		return nil, d.bizError(err)
	}
	return &m, nil
}

func (d *UserDao) SelectByOutUid(ctx context.Context, channelId, outUid string) (*po.UserInfoModel, error) {
	var m po.UserInfoModel
	err := d.db.Self.Where("channel_id = ? AND out_uid = ? AND user_status != 'deleted'", channelId, outUid).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		d.logger.Errorf("SelectByOutUid error: %v", err)
		return nil, d.bizError(err)
	}
	return &m, nil
}

func (d *UserDao) Insert(ctx context.Context, user *po.UserInfoModel) (*po.UserInfoModel, error) {
	err := d.db.Self.Create(user).Error
	if err != nil {
		d.logger.Errorf("Insert error: %v", err)
		return nil, d.bizError(err)
	}
	return user, nil
}

func (d *UserDao) Update(ctx context.Context, user *po.UserInfoModel) (*po.UserInfoModel, error) {
	err := d.db.Self.Omit("id, gmt_create").Select("*").Updates(user).Error
	if err != nil {
		d.logger.Errorf("Update error: %v", err)
		return nil, d.bizError(err)
	}
	return user, nil
}

func (d *UserDao) Upsert(ctx context.Context, user *po.UserInfoModel) (*po.UserInfoModel, error) {
	err := d.db.Self.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid"}},
		UpdateAll: true,
	}).Create(user).Error
	if err != nil {
		d.logger.Errorf("Upsert error: %v", err)
		return nil, d.bizError(err)
	}
	return user, nil
}
