package redis

import (
	"context"
	"time"

	redisV8 "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// ClientConfig ...
type ClientConfig struct {
	Network      string
	Addr         string
	Password     string
	DB           int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	PoolSize     int
}

// NewClient ...
func NewClient(ctx context.Context, config *ClientConfig) (*redisV8.Client, error) {
	options := &redisV8.Options{
		Network:      config.Network,
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  time.Duration(config.DialTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Millisecond,
		PoolSize:     config.PoolSize,
	}
	rds := redisV8.NewClient(options)
	_, err := rds.Ping(ctx).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rds, nil
}
