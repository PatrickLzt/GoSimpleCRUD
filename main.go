package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

var (
	posts   = make(map[int]Post)
	nextId  = 1
	postsMu sync.Mutex
)

func main() {

	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/", postHandler)

	fmt.Println("Server is listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
