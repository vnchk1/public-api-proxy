package main

import (
	"fmt"
	"github.com/vnchk1/public-api-proxy/client"
	"time"
)

func main() {
	baseURL := "https://jsonplaceholder.typicode.com/"
	timeout := 5 * time.Second
	postId := 1
	//инициализация клиента
	currentClient := client.NewRestyClient(baseURL, timeout)
	//запрос
	post, err := client.GetPostsRequest(currentClient, postId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Post: %v\n", post.ID)
	fmt.Printf("Title: %v\n", post.Title)
	fmt.Printf("Body: %v\n", post.Body)
	fmt.Printf("UserID: %v\n", post.UserID)
}
