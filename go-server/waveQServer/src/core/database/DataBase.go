package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
	"time"
	"waveQServer/src/utils/logutil"
)

var qdb *gorm.DB

// 初始化数据库链接
func init() {
	s := logutil.GetPath() + "lib" + string(filepath.Separator) + "qingluan.db"
	db, err := gorm.Open(sqlite.Open(s), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(15)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	qdb = db
}

// GetDb 获取数据库链接
func GetDb() *gorm.DB {
	return qdb
}
