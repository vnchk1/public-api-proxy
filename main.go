package main

import (
	"github.com/go-resty/resty/v2"
	"time"
)

func main() {
	defaultBaseURL := "https://jsonplaceholder.typicode.com/"
	client := resty.New()
	client.SetBaseURL(defaultBaseURL).
		SetTimeout(10*time.Second).
		SetHeader("Content-Type", "application/json")

}
