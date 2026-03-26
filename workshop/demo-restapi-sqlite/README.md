# Workshop REST API with SQLite

## Testing
```
$cd api
$go fmt ./...
$go test ./... -v
$go test ./... -v -coverpkg=./...
```

## Running
```
$cd api
$go run cmd/main.go
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

## Working with Docker
```
$docker compose up -d --build
$docker compose ps
```