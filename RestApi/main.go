package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				getUsers(w, r)
			} else if r.Method == http.MethodPost {
				createUser(w, r)
			}
		})

	http.HandleFunc("/users/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			getUserByID(w, r)
		})

	http.HandleFunc("/books",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				getBooks(w, r)
			} else if r.Method == http.MethodPost {
				createBook(w, r)
			}
		})

	http.HandleFunc("/users/{id}/books",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				addBookToUser(w, r)
			} else if r.Method == http.MethodGet {
				getUserBooks(w, r)
			}
		})

	http.HandleFunc(("/"),
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Welcome to the User-Book API ==> start server")
		})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
