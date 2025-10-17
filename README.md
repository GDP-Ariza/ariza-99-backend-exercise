# Backend Services Exercise

This repository contains two microservices built with Go and Gin framework:

- **User Service** - Manages user operations (port 7001)
- **Public Service** - Public-facing API that communicates with User Service (port 7002)

## Architecture

```
Client Request → Public Service (7002) → User Service (7001)
```

The Public Service acts as a gateway that forwards requests to the User Service and provides a unified public API.

## Prerequisites

- Go 1.24+ installed
- Git

## Quick Start

### 1. Clone the Repository

```bash
git clone <repository-url>
cd ariza-99-backend-exercise
```

### 2. Run User Service

```bash
cd user_service
go mod download
go run main.go
```

The User Service will start on **port 7001**.

### 3. Run Public Service

```bash
cd public_service
go mod download
go run main.go
```

The Public Service will start on **port 7002**.

## API Endpoints

### Public Service (Port 7002)

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| GET | `/ping` | Health check | - |
| GET | `/public-api/listings` | Get listings | Query params: `page_num`, `page_size`, `user_id` |
| POST | `/public-api/create-user` | Create new user | `{"name": "John Doe"}` |
| POST | `/public-api/create-listing` | Create new listing | `{"user_id": 1, "listing_type": "rent", "price": 5000}` |

### User Service (Port 7001)

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| GET | `/ping` | Health check | - |
| GET | `/users` | List users | Query params: `page_num`, `page_size` |
| POST | `/users` | Create user | Form data: `name=John Doe` |
| GET | `/users/{id}` | Get user by ID | - |

## Environment Variables

### Public Service

- `PORT` - Server port (default: 7002)
- `USER_SERVICE_URL` - User service URL (default: http://localhost:7001)
- `LISTING_SERVICE_URL` - Listing service URL (default: http://localhost:6000)

### User Service

- `PORT` - Server port (default: 7001)

## Example Usage

### Create a User via Public API

```bash
curl -X POST http://localhost:7002/public-api/create-user \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Johnson"}'
```

### Get Listings via Public API

```bash
curl "http://localhost:7002/public-api/listings?page_num=1&page_size=10"
```

### Create a Listing via Public API

```bash
curl -X POST http://localhost:7002/public-api/create-listing \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "listing_type": "rent", "price": 6000}'
```

## Development

### Project Structure

```
├── user_service/
│   ├── handler/          # HTTP handlers
│   ├── model/           # Data models
│   ├── repository/      # Data access layer
│   ├── service/         # Business logic
│   └── main.go
└── public_service/
    ├── adapter/         # External service adapters
    ├── handler/         # HTTP handlers
    ├── model/          # Data models
    ├── service/        # Business logic
    └── main.go
```

### Building

```bash
# Build user service
cd user_service
go build -o user_service

# Build public service
cd public_service
go build -o public_service
```

## Testing

### Health Check

```bash
# User service health check
curl http://localhost:7001/ping

# Public service health check
curl http://localhost:7002/ping
```

## Configuration

Both services support configuration through environment variables. For production deployments, consider using:

- Docker containers
- Kubernetes ConfigMaps/Secrets
- Environment-specific `.env` files

## Dependencies

- **Gin Web Framework** - HTTP router and middleware
- **Standard Go libraries** - HTTP client, JSON handling, etc.

## License

This project is part of a backend exercise.
