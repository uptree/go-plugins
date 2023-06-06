package clickhouse

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mailru/dbr"
	_ "github.com/mailru/go-clickhouse"
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
}

// NewConnection ...
func NewConnection(config *ClientConfig) *dbr.Connection {
	dsn := fmt.Sprintf("http://%s:%s@%s:%d/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	conn, err := dbr.Open("clickhouse", dsn, nil)
	if err != nil {
		panic(err)
	}
	conn.SetMaxIdleConns(config.MaxIdleConns)
	conn.SetMaxOpenConns(config.MaxOpenConns)
	conn.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute)
	conn.SetConnMaxIdleTime(time.Duration(config.ConnMaxIdleTime) * time.Minute)
	return conn
}

// NewClient ...
func NewClient(config *ClientConfig) *sql.DB {
	return NewConnection(config).DB
}
