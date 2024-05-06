package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConf DbConfig

type DbConfig struct {
	Host        string
	Port        uint
	DbName      string
	User        string
	Password    string
	DBZone      string
	Charset     string
	MaxIdle     int
	MaxOpen     int
	LogMode     bool
	Loc         string
	MaxLifetime int64
	TablePrefix string
	Debug       bool
}

func NewDbClient(c *DbConfig) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf(format,
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
	)

	// 处理表前缀
	var newLogger logger.Interface
	newLogger = NewGormZapLogger()
	if c.Debug {
		newLogger = newLogger.LogMode(logger.Info)
	} else {
		newLogger = newLogger.LogMode(logger.Silent)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, sqlErr := db.DB()
	if c.MaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(c.MaxLifetime))
	} else {
		sqlDB.SetConnMaxLifetime(time.Second * 1000)
	}
	sqlDB.SetMaxOpenConns(c.MaxOpen)
	sqlDB.SetMaxIdleConns(c.MaxIdle)

	if sqlErr != nil {
		return nil, sqlErr
	}

	// Ping
	if pingErr := sqlDB.Ping(); pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}

func GetDBConf() *DbConfig {
	return &DBConf
}
