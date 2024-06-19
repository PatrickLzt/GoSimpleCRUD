package main

import (
	"net/http"
	"strconv"
)

func postsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		handleGetPosts(w, r)

	case "POST":
		handlePostPosts(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.Error(w, "Invalid post id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		handleGetPost(w, r, id)

	case "DELETE":
		handleDeletePost(w, r, id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
