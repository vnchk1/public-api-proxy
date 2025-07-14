package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vnchk1/public-api-proxy/client"
	"github.com/vnchk1/public-api-proxy/configs"
	"log"
	"os"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatalf("failed to load .env file")
	}

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	cfg, err := configs.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	//инициализация клиента
	newClient := client.NewRestyClient(cfg)
	//запрос
	post, err := client.GetPostsRequest(newClient, cfg.PostId)
	if err != nil {
		log.Fatalf("failed to get posts: %v", err)
	}

	fmt.Printf("Post: %v\n", post.ID)
	fmt.Printf("Title: %v\n", post.Title)
	fmt.Printf("Body: %v\n", post.Body)
	fmt.Printf("UserID: %v\n", post.UserID)
}
