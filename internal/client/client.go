package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/vnchk1/public-api-proxy/internal/config"
	"log/slog"
	"time"
)

const (
	TimeoutSeconds          = 5
	RetryCount              = 3
	RetryWaitTimeSeconds    = 2
	RetryMaxWaitTimeSeconds = 10
)

func NewRestyClient(cfg *config.Config, logger *slog.Logger) (client *resty.Client) {
	client = resty.New()
	client.SetBaseURL(cfg.BaseUrl).
		SetTimeout(TimeoutSeconds*time.Second).
		SetHeader("Content-Type", "application/json").
		SetRetryCount(RetryCount).
		SetRetryWaitTime(RetryWaitTimeSeconds * time.Second).
		SetRetryMaxWaitTime(RetryMaxWaitTimeSeconds * time.Second)

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		logger.Info("response received",
			"status code", resp.StatusCode(),
			"url", resp.Request.URL,
			"method", resp.Request.Method,
			"time", resp.Request.Time)
		return nil
	})
	return
}

func GetPostsRequest(client *resty.Client, id string) (post *Post, err error) {
	postPath := fmt.Sprintf("posts/%v", id)
	resp, err := client.R().SetResult(&post).Get(postPath)

	if err != nil {
		return nil, fmt.Errorf("sending request error: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API error: %v %v", resp.Error(), resp.StatusCode())
	}

	return
}
