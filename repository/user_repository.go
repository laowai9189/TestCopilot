package repository

import (
	"fmt"
	"log"

	"github.com/laowai9189/TestCopilot/models"
)

// UserRepository handles user database operations
type UserRepository struct {
	// In production, this would contain database connection
}

// mockUsers simulates database storage for demonstration
var mockUsers = []models.User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
}
var nextID = 3

// NewUserRepository creates a new user repository instance
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetAll retrieves all users from the database
func (r *UserRepository) GetAll() ([]models.User, error) {
	log.Printf("Mock SQL: SELECT * FROM users")
	return mockUsers, nil
}

// GetByID retrieves a user by ID from the database
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	log.Printf("Mock SQL: SELECT * FROM users WHERE id = %d", id)
	
	for _, user := range mockUsers {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

// Create creates a new user in the database
func (r *UserRepository) Create(user *models.User) error {
	log.Printf("Mock SQL: INSERT INTO users (name, email) VALUES ('%s', '%s')", user.Name, user.Email)
	
	// Assign ID and add to mock storage
	user.ID = nextID
	nextID++
	mockUsers = append(mockUsers, *user)
	
	log.Printf("User created with ID: %d", user.ID)
	return nil
}

// Update updates an existing user in the database
func (r *UserRepository) Update(id int, user *models.User) error {
	log.Printf("Mock SQL: UPDATE users SET name = '%s', email = '%s' WHERE id = %d", user.Name, user.Email, id)
	
	for i, existingUser := range mockUsers {
		if existingUser.ID == id {
			user.ID = id
			mockUsers[i] = *user
			log.Printf("User with ID %d updated successfully", id)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

// Delete deletes a user from the database
func (r *UserRepository) Delete(id int) error {
	log.Printf("Mock SQL: DELETE FROM users WHERE id = %d", id)
	
	for i, user := range mockUsers {
		if user.ID == id {
			mockUsers = append(mockUsers[:i], mockUsers[i+1:]...)
			log.Printf("User with ID %d deleted successfully", id)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}