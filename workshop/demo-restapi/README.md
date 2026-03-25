# Demo REST API with Go
* Gin web framework
* PostgreSQL database
* Redis cache
* Docker Compose for local development


## Setup environment variables
1. Create a `.env` file in the project root directory.
2. Copy the contents of `.env.sample` into `.env`.
3. Update the values in `.env` if necessary (e.g., database credentials).

## Run the application
```
$docker compose up --build
$docker compose ps
```

## API Endpoints

Create a new product:
```
curl -X POST http://localhost:8080/products \
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