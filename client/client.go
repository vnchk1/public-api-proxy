package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/vnchk1/public-api-proxy/internal/models"
	"time"
)

func NewRestyClient(baseUrl string) *resty.Client {
	client := resty.New()
	client.SetBaseURL(baseUrl).
		SetTimeout(5*time.Second).
		SetHeader("Content-Type", "application/json").
		SetRetryCount(3).
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second)
	return client
}

func GetPostsRequest(client *resty.Client, id string) (*models.Post, error) {
	var Post *models.Post
	postPath := fmt.Sprintf("posts/%v", id)
	resp, err := client.R().SetResult(&Post).Get(postPath)

	if err != nil {
		return nil, fmt.Errorf("sending request error: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API error: %v %v", resp.Error(), resp.StatusCode())
	}
	return Post, err
}
