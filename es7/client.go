package es7

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

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
	if config.LogFile == nil {
		config.LogFile = os.Stdout
	}
	esClient, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.URL),
		elastic.SetBasicAuth(config.Username, config.Password),
		elastic.SetHttpClient(ElasticV7Doer),
		elastic.SetInfoLog(log.New(config.LogFile, "[ES-INFO]", log.LstdFlags)),
		elastic.SetErrorLog(log.New(config.LogFile, "[ES-ERROR]", log.LstdFlags)),
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
