package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/vnchk1/public-api-proxy/configs"
	"github.com/vnchk1/public-api-proxy/internal/models"
	"github.com/vnchk1/public-api-proxy/logging"
	"time"
)

const (
	TimeoutSeconds   = 5
	RetryCount       = 3
	RetryWaitTime    = 2
	RetryMaxWaitTime = 10
)

func NewRestyClient(cfg *configs.Configs) *resty.Client {
	client := resty.New()
	client.SetBaseURL(cfg.BaseUrl).
		SetTimeout(TimeoutSeconds*time.Second).
		SetHeader("Content-Type", "application/json").
		SetRetryCount(RetryCount).
		SetRetryWaitTime(RetryWaitTime * time.Second).
		SetRetryMaxWaitTime(RetryMaxWaitTime * time.Second)

	logger := logging.NewLogger(cfg.LogLevel)
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		logger.Info("response received",
			"status code", resp.StatusCode(),
			"url", resp.Request.URL,
			"method", resp.Request.Method,
			"time", resp.Request.Time)
		return nil
	})
	return client
}

func GetPostsRequest(client *resty.Client, id string) (post *models.Post, err error) {
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
