package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/laowai9189/TestCopilot/handlers"
	"github.com/laowai9189/TestCopilot/models"
)

func main() {
	// Create router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()
	
	// Health check
	api.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	
	// User routes
	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	// Root endpoint
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := models.Response{
			Status:  "success",
			Message: "Welcome to TestCopilot Backend API",
			Data: map[string]interface{}{
				"version": "1.0.0",
				"endpoints": []string{
					"GET /api/v1/health",
					"GET /api/v1/users",
					"POST /api/v1/users",
					"GET /api/v1/users/{id}",
					"PUT /api/v1/users/{id}",
					"DELETE /api/v1/users/{id}",
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	// Start server
	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	fmt.Printf("Available endpoints:\n")
	fmt.Printf("  GET  /                    - API information\n")
	fmt.Printf("  GET  /api/v1/health       - Health check\n")
	fmt.Printf("  GET  /api/v1/users        - Get all users\n")
	fmt.Printf("  POST /api/v1/users        - Create user\n")
	fmt.Printf("  GET  /api/v1/users/{id}   - Get user by ID\n")
	fmt.Printf("  PUT  /api/v1/users/{id}   - Update user\n")
	fmt.Printf("  DELETE /api/v1/users/{id} - Delete user\n")
	
	log.Fatal(http.ListenAndServe(":"+port, r))
}