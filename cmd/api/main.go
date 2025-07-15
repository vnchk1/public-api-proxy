package main

import (
	"github.com/joho/godotenv"
	"github.com/vnchk1/public-api-proxy/client"
	"github.com/vnchk1/public-api-proxy/configs"
	"github.com/vnchk1/public-api-proxy/logging"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
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

	logger := logging.NewLogger(cfg.LogLevel)
	//инициализация клиента
	newClient := client.NewRestyClient(cfg)
	//запрос
	post, err := client.GetPostsRequest(newClient, cfg.PostId)
	if err != nil {
		log.Fatalf("failed to get posts: %v", err)
	}

	logger.Info("response result",
		"Post ID", post.ID,
		"Title", post.Title,
		"Body", post.Body,
		"User ID", post.UserID)
}
