package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/laowai9189/TestCopilot/models"
	"github.com/laowai9189/TestCopilot/repository"
)

// userRepo is the user repository instance
var userRepo *repository.UserRepository

func init() {
	// Initialize the user repository
	userRepo = repository.NewUserRepository()
}

// HealthCheck handles health check requests
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Status:  "success",
		Message: "Server is running",
		Data: map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
			"service":   "TestCopilot Backend",
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUsers returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userRepo.GetAll()
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: "Failed to retrieve users",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    users,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUser returns a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := userRepo.GetByID(id)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: "User not found",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "User found",
		Data:    user,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: "Invalid JSON payload",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = userRepo.Create(&user)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: "Failed to create user",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "User created successfully",
		Data:    user,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// UpdateUser updates an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: "Invalid JSON payload",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = userRepo.Update(id, &updatedUser)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: "User not found",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "User updated successfully",
		Data:    updatedUser,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteUser deletes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = userRepo.Delete(id)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: "User not found",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "User deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}