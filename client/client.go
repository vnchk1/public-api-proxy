package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BaseURL = "https://jsonplaceholder.typicode.com/"
)

type Post struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"userId"`
}

func GetPost(ctx context.Context, id int64) (*Post, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseURL+fmt.Sprintf("posts/%v", id), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var post Post
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return &post, nil
}
