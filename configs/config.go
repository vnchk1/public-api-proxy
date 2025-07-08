package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
)

type Configs struct {
	BaseUrl string `yaml:"base_url"`
	PostId  string `yaml:"post_id"`
}

func LoadConfig(path string) (*Configs, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file")
	}
	config := &Configs{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	if envBaseUrl := os.Getenv("BASE_URL"); envBaseUrl != "" {
		config.BaseUrl = envBaseUrl
	}
	if envPostId := os.Getenv("POST_ID"); envPostId != "" {
		config.PostId = envPostId
	}

	return config, nil
}
