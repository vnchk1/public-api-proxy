package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Post struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"userId"`
}

func main() {
	baseURL := "https://jsonplaceholder.typicode.com/"
	client := resty.New()
	client.SetBaseURL(baseURL).
		SetTimeout(10*time.Second).
		SetHeader("Content-Type", "application/json")
	//get-запрос к posts/{id}
	var post Post
	postId := "1"
	postPath := fmt.Sprintf("posts/%s", postId)
	resp, err := client.R().SetResult(&post).Get(postPath)

	if err != nil {
		fmt.Printf("Sending request error: %v", err)
		return
	}

	if resp.IsError() {
		fmt.Printf("API error: %v %v", resp.Error(), resp.StatusCode())
	}

	fmt.Printf("Post: %v\n", post.ID)
	fmt.Printf("Title: %v\n", post.Title)
	fmt.Printf("Body: %v\n", post.Body)
	fmt.Printf("UserID: %v\n", post.UserID)
}
