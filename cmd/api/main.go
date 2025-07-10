package main

import (
	"context"
	"fmt"
	"github.com/vnchk1/public-api-proxy/client"
	"log"
)

const PostId = 1

func main() {

	ctx := context.Background()

	postResp, err := client.GetPost(ctx, PostId)
	if err != nil {
		log.Fatalf("Getting posts/ error:%v", err)
	}

	fmt.Printf("Post ID: %v\n", postResp.ID)
	fmt.Printf("Title: %v\n", postResp.Title)
	fmt.Printf("Body: %v\n", postResp.Body)
	fmt.Printf("User ID: %v\n", postResp.UserID)
}
