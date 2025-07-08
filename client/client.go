package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/vnchk1/public-api-proxy/internal/models"
	"time"
)

func NewRestyClient(baseURL string, timeout time.Duration) *resty.Client {
	client := resty.New()
	client.SetBaseURL(baseURL).
		SetTimeout(timeout).
		SetHeader("Content-Type", "application/json")
	return client
}

func GetPostsRequest(client *resty.Client, id int) (*models.Post, error) {
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
