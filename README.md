# TestCopilot

A simple Golang backend API project built with [gorilla/mux](https://github.com/gorilla/mux) framework.

## Features

- RESTful API with CRUD operations for users
- JSON responses with standardized format
- Health check endpoint
- In-memory data storage for demonstration
- Clean project structure with separated handlers and models

## Project Structure

```
.
├── main.go           # Main application entry point
├── go.mod           # Go module dependencies
├── go.sum           # Go module checksums
├── handlers/        # HTTP request handlers
│   └── user.go      # User-related handlers
├── models/          # Data models
│   └── user.go      # User model and response structures
└── README.md        # This file
```

## Prerequisites

- Go 1.24+ installed on your system
- Git (for cloning the repository)

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

The server will start on port 8080 by default.

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
2. Create handlers in `handlers/` directory
3. Register routes in `main.go`

### Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router and URL matcher

## License

This project is open source and available under the [MIT License](LICENSE).