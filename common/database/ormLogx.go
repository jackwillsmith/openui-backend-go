package database

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"time"
)

type dbLogx struct {
	LogLevel                  gormLogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func (d dbLogx) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return dbLogx{
		SlowThreshold:             d.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          d.SkipCallerLookup,
		IgnoreRecordNotFoundError: d.IgnoreRecordNotFoundError,
	}
}

func (d dbLogx) Info(ctx context.Context, s string, i ...interface{}) {
	if d.LogLevel < gormLogger.Info {
		return
	}
	logx.WithContext(ctx).Infof(s, i...)
}

func (d dbLogx) Warn(ctx context.Context, s string, i ...interface{}) {
	if d.LogLevel < gormLogger.Warn {
		return
	}
	logx.WithContext(ctx).Slowf(s, i...)
}

func (d dbLogx) Error(ctx context.Context, s string, i ...interface{}) {
	if d.LogLevel < gormLogger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(s, i...)
}

func (d dbLogx) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if d.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && d.LogLevel >= gormLogger.Error && (!d.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		logx.WithContext(ctx).WithDuration(elapsed).Errorf("err: %s, sql: %s, rows: %d", err, sql, rows)
	case d.SlowThreshold != 0 && elapsed > d.SlowThreshold && d.LogLevel >= gormLogger.Warn:
		sql, rows := fc()
		logx.WithContext(ctx).WithDuration(elapsed).Slowf("sql: %s, rows: %d", sql, rows)
	case d.LogLevel >= gormLogger.Info:
		sql, rows := fc()
		logx.WithContext(ctx).WithDuration(elapsed).Infof("sql: %s, rows: %d", sql, rows)
	}
}

func NewGormZapLogger() gormLogger.Interface {
	return dbLogx{
		LogLevel:                  gormLogger.Info,
		SlowThreshold:             time.Second,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}
