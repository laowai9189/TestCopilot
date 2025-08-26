package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/laowai9189/TestCopilot/config"
)

var DB *sql.DB

// Initialize initializes the database connection
func Initialize() error {
	dbConfig := config.GetDatabaseConfig()
	
	// Print the generated configuration for demonstration
	fmt.Printf("Database Configuration:\n")
	fmt.Printf("  Host: %s\n", dbConfig.Host)
	fmt.Printf("  Port: %d\n", dbConfig.Port)
	fmt.Printf("  Username: %s\n", dbConfig.Username)
	fmt.Printf("  Password: %s\n", dbConfig.Password)
	fmt.Printf("  Database: %s\n", dbConfig.Database)
	fmt.Printf("  DSN: %s\n", dbConfig.GetDSN())
	
	// For demonstration purposes, we'll use a mock connection
	// In a real scenario, you would connect to an actual MySQL database
	fmt.Printf("Note: This is a demonstration. In production, you would connect to a real MySQL database.\n")
	
	// Create a mock database connection for demonstration
	// In production, you would use: db, err := sql.Open("mysql", dbConfig.GetDSN())
	return createMockConnection()
}

// createMockConnection creates a mock database connection for demonstration
func createMockConnection() error {
	// For demonstration, we'll simulate a successful connection
	log.Println("Mock MySQL database connection established successfully")
	return nil
}

// CreateUsersTable creates the users table in the database
func CreateUsersTable() error {
	// In a real MySQL implementation, this would create the actual table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`
	
	fmt.Printf("Mock SQL executed: %s\n", createTableSQL)
	log.Println("Users table created successfully (mock)")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}