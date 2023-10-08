package es7

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

// ClientConfig ...
type ClientConfig struct {
	URL      string
	Username string
	Password string
	LogFile  io.Writer
}

// ElasticV7Doer ...
var ElasticV7Doer elastic.Doer = http.DefaultClient

// NewClient ...
func NewClient(ctx context.Context, config *ClientConfig) (*elastic.Client, error) {
	esClient, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.URL),
		elastic.SetBasicAuth(config.Username, config.Password),
		elastic.SetHttpClient(ElasticV7Doer),
		elastic.SetInfoLog(log.New(config.LogFile, "ES-INFO", 0)),
		elastic.SetTraceLog(log.New(config.LogFile, "ES-TRACE", 0)),
		elastic.SetErrorLog(log.New(config.LogFile, "ES-ERROR", 0)),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	_, _, err = esClient.Ping(config.URL).Do(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return esClient, nil
}
