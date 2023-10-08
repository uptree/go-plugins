package ck

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ClientConfig ...
type ClientConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	ConnMaxIdleTime int
	Timeout         string
	Logger          logger.Interface
}

// NewClient ...
func NewClient(ctx context.Context, config *ClientConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%d/%s?dial_timeout=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.Timeout,
	)
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	connMaxLifetime := time.Duration(config.ConnMaxLifetime)
	connMaxIdleTime := time.Duration(config.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(connMaxLifetime * time.Minute)
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime * time.Minute)
	err = sqlDB.PingContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
