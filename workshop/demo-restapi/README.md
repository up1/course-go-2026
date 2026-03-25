# Demo REST API with Go
* Gin web framework
* PostgreSQL database
* Redis cache
* Docker Compose for local development

## Setup environment variables
1. Create a `.env` file in the project root directory.
2. Copy the contents of `.env.sample` into `.env`.
3. Update the values in `.env` if necessary (e.g., database credentials).

## Run with go
```
# Install dependencies and format code
$go mod tidy
$go fmt ./...
$go fix ./...

# Testing with caching (may return cached results)
$go test ./... -v
$go test ./... -count=1 -v

# Testing with coverage report
$go test -v -coverpkg=./... -coverprofile=coverage.out ./...
$go tool cover -func=coverage.out
$go tool cover -html=coverage.out
```

## Run the application with Docker Compose
```
$docker compose down
$docker compose up -d --build
$docker compose ps
```

## API Endpoints

Create a new product:
```
$curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{
    "name": "Product 1",
    "description": "This is product 1",
    "price": 9.99
}'  
```

Get product by ID:
```
$curl  http://localhost:8080/products/1
```