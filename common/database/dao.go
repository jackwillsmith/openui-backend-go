package database

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
)

/**
* Author: chenlc
* TODO: test
* Date: 2024/3/27
**/

// GormDao gorm dao.
type GormDao struct {
	db *gorm.DB
}

// New a gorm dao and return.
func New(config DbConfig) (dao *GormDao) {
	dao = &GormDao{}

	// note: 添加数据库的链接
	db, err := NewDbClient(&config)
	if err != nil {
		logc.Error(context.Background(), logc.Field("connect mysql db error", err))
		return nil
	}
	dao.db = db
	return
}

// Close the resource.
func (d *GormDao) Close() {
	dbbase, _ := d.db.DB()
	dbbase.Close()
}

// Ping ping the resource.
func (d *GormDao) Ping(ctx context.Context) (err error) {
	// note: 设置5秒的超时
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	dbbase, _ := d.db.DB()
	defer cancel()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return dbbase.Ping()
	}
}

// GetDB get the resource.
func (d *GormDao) GetDB() *gorm.DB {
	return d.db
}
