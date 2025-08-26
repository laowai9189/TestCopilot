package config

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// DatabaseConfig holds the MySQL database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

// generateRandomPort generates a random port between 3000 and 9999
func generateRandomPort() int {
	num, _ := rand.Int(rand.Reader, big.NewInt(6999))
	return int(num.Int64()) + 3000
}

// generateRandomIP generates a random IP address (192.168.x.x range)
func generateRandomIP() string {
	num1, _ := rand.Int(rand.Reader, big.NewInt(255))
	num2, _ := rand.Int(rand.Reader, big.NewInt(255))
	return fmt.Sprintf("192.168.%d.%d", num1.Int64(), num2.Int64())
}

// GetDatabaseConfig returns a randomly generated database configuration
func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     generateRandomIP(),
		Port:     generateRandomPort(),
		Username: generateRandomString(8),
		Password: generateRandomString(12),
		Database: "testcopilot_db",
	}
}

// GetDSN returns the Data Source Name for MySQL connection
func (config *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)
}