package mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
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
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.Timeout,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 config.Logger,
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
	})
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
	sqlDB.SetConnMaxLifetime(connMaxLifetime * time.Minute)
	err = sqlDB.PingContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
