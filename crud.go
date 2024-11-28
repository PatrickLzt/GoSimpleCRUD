package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CRUD
func handleGetUsers(w http.ResponseWriter, r *http.Request) {

	usersMu.Lock()
	defer usersMu.Unlock()

	ps := make([]User, 0, len(users))

	for _, u := range users {
		ps = append(ps, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ps)

	fmt.Println("GET /users", ps)

}

func handlePostUsers(w http.ResponseWriter, r *http.Request) {

	var u User

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &u); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	usersMu.Lock()
	defer usersMu.Unlock()

	u.Id = int64(nextId)
	nextId++
	users[int(u.Id)] = u

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Println("POST /users", u)
}

func handleSingleGetUser(w http.ResponseWriter, r *http.Request, id int) {

	usersMu.Lock()
	defer usersMu.Unlock()

	u, ok := users[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)

	fmt.Println("GET /users/", id, u)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, id int) {

	usersMu.Lock()
	defer usersMu.Unlock()

	if _, ok := users[id]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, id)
	w.WriteHeader(http.StatusNoContent)

	fmt.Println("DELETE /users/", id)
}
