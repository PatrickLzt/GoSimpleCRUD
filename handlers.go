package main

import (
	"net/http"
	"strconv"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		handleGetUsers(w, r)

	case "POST":
		handlePostUsers(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func singleUserHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		http.Error(w, "Invalid post id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		handleSingleGetUser(w, r, id)

	case "DELETE":
		handleDeleteUser(w, r, id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
