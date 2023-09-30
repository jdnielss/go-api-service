package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User is a struct to represent user data
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Response is a custom struct for the API response
type Response struct {
	Status struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"status"`
	Data any `json:"data,omitempty"`
}

func main() {
	log.Println("Starting the HTTP server on port 8080")

	http.HandleFunc("/", index)
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/users", GetUsers)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Health Check")

	response := Response{
		Status: struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{
			Message: "OK",
			Code:    http.StatusOK,
		},
		Data: "Hello World",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Health Check")

	response := Response{
		Status: struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{
			Message: "OK",
			Code:    http.StatusOK,
		},
		Data: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Users")

	// Simulate fetching user data
	users := []User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}

	// Create a Response instance
	response := Response{
		Status: struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{
			Message: "OK",
			Code:    http.StatusOK,
		},
		Data: users,
	}

	// Encode the Response as JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
