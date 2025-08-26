# TestCopilot

A Golang backend API project built with [gorilla/mux](https://github.com/gorilla/mux) framework and MySQL database integration.

## Features

- RESTful API with CRUD operations for users
- JSON responses with standardized format
- Health check endpoint
- MySQL database integration with mock implementation for demonstration
- Clean project structure with separated handlers, models, and repository patterns
- Randomly generated MySQL configuration for development

## Project Structure

```
.
├── main.go           # Main application entry point
├── go.mod           # Go module dependencies
├── go.sum           # Go module checksums
├── config/          # Configuration packages
│   └── database.go  # Database configuration with random generation
├── database/        # Database connection and management
│   └── connection.go # MySQL connection handling
├── handlers/        # HTTP request handlers
│   └── user.go      # User-related handlers
├── models/          # Data models
│   └── user.go      # User model and response structures
├── repository/      # Data access layer
│   └── user_repository.go # User database operations
└── README.md        # This file
```

## Prerequisites

- Go 1.24+ installed on your system
- Git (for cloning the repository)
- MySQL database (for production use)

## Database Configuration

The application automatically generates random MySQL connection details for demonstration:
- Random IP address (192.168.x.x range)
- Random port (3000-9999)
- Random username and password
- Database name: `testcopilot_db`

**Note**: This implementation includes a mock database layer for demonstration purposes. In production, you would connect to a real MySQL database by updating the connection logic in `database/connection.go`.

## Installation & Setup

1. Clone the repository:
```bash
git clone https://github.com/laowai9189/TestCopilot.git
cd TestCopilot
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o testcopilot .
```

4. Run the server:
```bash
./testcopilot
```

Or run directly with Go:
```bash
go run main.go
```

The server will start on port 8080 by default and display the randomly generated database configuration.

## API Endpoints

### Root
- `GET /` - API information and available endpoints

### Health Check
- `GET /api/v1/health` - Server health status

### Users
- `GET /api/v1/users` - Get all users
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users/{id}` - Get user by ID
- `PUT /api/v1/users/{id}` - Update user by ID
- `DELETE /api/v1/users/{id}` - Delete user by ID

## API Usage Examples

### Get API Information
```bash
curl http://localhost:8080/
```

### Health Check
```bash
curl http://localhost:8080/api/v1/health
```

### Get All Users
```bash
curl http://localhost:8080/api/v1/users
```

### Create a New User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Johnson", "email": "alice@example.com"}'
```

### Get User by ID
```bash
curl http://localhost:8080/api/v1/users/1
```

### Update User
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "John Updated", "email": "john.updated@example.com"}'
```

### Delete User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Response Format

All API responses follow a standard JSON format:

```json
{
  "status": "success|error",
  "message": "Description of the result",
  "data": "Optional data payload"
}
```

## Data Model

### User
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com"
}
```

## Development

### Adding New Endpoints

1. Define new models in `models/` directory
2. Create repository methods in `repository/` directory
3. Create handlers in `handlers/` directory
4. Register routes in `main.go`

### Database Integration

The project uses a repository pattern for database operations:
- `config/database.go` - Database configuration management
- `database/connection.go` - Connection handling and initialization
- `repository/user_repository.go` - User data access operations

To integrate with a real MySQL database:
1. Update the connection logic in `database/connection.go`
2. Replace mock operations in `repository/user_repository.go` with actual SQL queries
3. Configure your MySQL database connection details

### Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router and URL matcher
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go

## License

This project is open source and available under the [MIT License](LICENSE).