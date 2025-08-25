package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler for the root path → serve static HTML
func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

// Handler for /hello → return plain text
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go Web Server!")
}

// Handler for /api → return JSON
func apiHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello, JSON!",
		"status":  "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Routes
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/api", apiHandler)

	// Start server on port 8080
	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
