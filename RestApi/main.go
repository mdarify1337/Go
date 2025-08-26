package main

import (
	"log"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow your frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getUsers(w, r)
		} else if r.Method == http.MethodPost {
			createUser(w, r)
		}
	})

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		getUserByID(w, r)
	})

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getBooks(w, r)
		} else if r.Method == http.MethodPost {
			createBook(w, r)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			addBookToUser(w, r)
		} else if r.Method == http.MethodGet {
			getUserBooks(w, r)
		}
	})

	// Wrap mux with CORS middleware
	handler := enableCORS(mux)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
