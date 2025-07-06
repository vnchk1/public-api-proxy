package main

import (
	"context"
	"fmt"
	"github.com/vnchk1/public-api-proxy/client"
	"net/http"
	"time"
)

func main() {
	httpClient := client.NewHTTPClient(&http.Client{
		Timeout: 10 * time.Second,
	})

	ctx := context.Background()

	Post, _, err := httpClient.GetPost(ctx, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Post:")
	fmt.Printf("ID: %v\n", Post.ID)
	fmt.Printf("Title: %v\n", Post.Title)
	fmt.Printf("Body: %v\n", Post.Body)
	fmt.Printf("User ID: %v\n", Post.UserID)
}
