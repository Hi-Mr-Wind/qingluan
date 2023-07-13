package dao

import (
	log "github.com/sirupsen/logrus"
	"go-service/pkg/errno"
	"go-service/pkg/repository"
)

type baseDao struct {
	db     *repository.Database
	logger *log.Entry
}

func newBaseDao(repo *repository.Database) *baseDao {
	return &baseDao{
		db:     repo,
		logger: log.WithField("type", "database"),
	}
}

func (d *baseDao) bizError(err error) error {
	return errno.NewBizError(errno.ErrDatabase, err)
}
