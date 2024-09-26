package http

import (
	"gihub.com/robbitancor/simple-microservice/internal/config"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

type Client struct {
	BaseUri string
	c       *http.Client
}

func NewClient(cfg config.ApiConfig, client *http.Client) *Client {
	return &Client{
		cfg.BaseUri,
		client,
	}
}

func (c Client) GetRequest(uri string) ([]byte, error) {
	resp, err := c.c.Get(uri)

	if err != nil {
		log.Error(err)
	}

	val, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
	}

	return val, err
}
