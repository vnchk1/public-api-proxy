package main

import (
	"fmt"
	"github.com/vnchk1/public-api-proxy/client"
	"github.com/vnchk1/public-api-proxy/configs"
	"log"
)

func main() {
	cfg, err := configs.LoadConfig(".yml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	//инициализация клиента
	newClient := client.NewRestyClient(cfg.BaseUrl)
	//запрос
	post, err := client.GetPostsRequest(newClient, cfg.PostId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Post: %v\n", post.ID)
	fmt.Printf("Title: %v\n", post.Title)
	fmt.Printf("Body: %v\n", post.Body)
	fmt.Printf("UserID: %v\n", post.UserID)
}
