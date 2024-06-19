package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CRUD
func handleGetPosts(w http.ResponseWriter, r *http.Request) {

	postsMu.Lock()
	defer postsMu.Unlock()

	ps := make([]Post, 0, len(posts))

	for _, p := range posts {
		ps = append(ps, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ps)

	fmt.Println("GET /posts", ps)

}

func handlePostPosts(w http.ResponseWriter, r *http.Request) {

	var p Post

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	postsMu.Lock()
	defer postsMu.Unlock()

	p.Id = nextId
	nextId++
	posts[p.Id] = p

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Println("POST /posts", p)
}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {

	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := posts[id]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

	fmt.Println("GET /posts/", id, p)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {

	postsMu.Lock()
	defer postsMu.Unlock()

	if _, ok := posts[id]; !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	delete(posts, id)
	w.WriteHeader(http.StatusNoContent)

	fmt.Println("DELETE /posts/", id)
}
