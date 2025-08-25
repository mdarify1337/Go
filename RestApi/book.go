package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	UserID int    `json:"user_id"` // Foreign key → belongs to a User
}

var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	// http.NotFound(w, r)
}

// Create a book for a user
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = len(books) + 1
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func addBookToUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["id"])

	// Find user
	for i, user := range users {
		if user.ID == userID {
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			bookIDCounter := 0
			book.ID = bookIDCounter
			bookIDCounter++
			book.UserID = userID

			books = append(books, book)
			users[i].Books = append(users[i].Books, book)

			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

func getUserBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["id"])

	for _, user := range users {
		if user.ID == userID {
			json.NewEncoder(w).Encode(user.Books)
			return
		}
	}
	http.NotFound(w, r)
}
