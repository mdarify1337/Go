package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Books []Book `json:"books,omitempty"` // A user can have many books
}

var users []User

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	fmt.Println("Get all users called", users)
	// http.NotFound(w, r)
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	fmt.Println("Create user called", user)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	fmt.Println("Get user by ID called with ID:", id)
	// http.NotFound(w, r)
}
