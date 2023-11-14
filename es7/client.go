package es7

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
)

// ClientConfig ...
type ClientConfig struct {
	URL      string
	Username string
	Password string
	LogFile  io.Writer
	LogLevel string
}

// ElasticV7Doer ...
var ElasticV7Doer elastic.Doer = http.DefaultClient

// NewClient ...
func NewClient(ctx context.Context, config *ClientConfig) (*elastic.Client, error) {
	var elsOpts []elastic.ClientOptionFunc
	elsOpts = append(elsOpts,
		elastic.SetSniff(false),
		elastic.SetURL(config.URL),
		elastic.SetBasicAuth(config.Username, config.Password),
		elastic.SetHttpClient(ElasticV7Doer),
	)
	if config.LogLevel != "" {
		if config.LogFile == nil {
			config.LogFile = os.Stdout
		}
		var (
			errorLog = elastic.SetErrorLog(log.New(config.LogFile, "[ES-ERROR] ", log.LstdFlags))
			infoLog  = elastic.SetInfoLog(log.New(config.LogFile, "[ES-INFO] ", log.LstdFlags))
			traceLog = elastic.SetTraceLog(log.New(config.LogFile, "[ES-TRACE] ", log.LstdFlags))
		)
		switch config.LogLevel {
		case "error":
			elsOpts = append(elsOpts, errorLog)
		case "info":
			elsOpts = append(elsOpts, errorLog, infoLog)
		case "trace":
			elsOpts = append(elsOpts, errorLog, infoLog, traceLog)
		}
	}
	esClient, err := elastic.NewClient(elsOpts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	_, _, err = esClient.Ping(config.URL).Do(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return esClient, nil
}
