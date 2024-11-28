package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	users   = make(map[int]User)
	nextId  = 1
	usersMu sync.Mutex
)

func main() {

	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", singleUserHandler)

	fmt.Println("Server is listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
