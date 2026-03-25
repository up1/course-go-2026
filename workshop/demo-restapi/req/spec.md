# Develop REST API with Golang and Gin framework

## Technical Requirements
1. Use Golang 1.26.1 as the programming language.
2. Use the Gin framework version v1.12.0 for building the REST API.
3. Implement the following endpoints:
   - `POST /products`: Create a new product.
   - `GET /products/{id}`: Retrieve a product by its ID.
4. Use an postgreSQL database to store product information and caching with Redis.
5. Implement input validation for the product creation endpoint.
6. Manage error handling and return appropriate HTTP status codes and error messages for different scenarios (e.g., product not found, validation errors, database errors). And use best practices for error handling in Go.
7. Handle errors gracefully and return appropriate HTTP status codes and error messages.
8. Write unit tests for the API endpoints.
9. Use environment variables for configuration (e.g., database connection strings, Redis connection details).
10. Follow best practices for structuring a Go project, including separation of concerns and modular design
11. Use dependency injection to manage dependencies between different components of the application.
12. Dockerize the application for easy deployment with docker compose.

## Project Structure
* User domain based on clean architecture principles.
* Separate layers for handlers, services, repositories, and models.
* Use dependency injection to manage dependencies between layers.

Sample project structure:
```
├── cmd
│   └── main.go
├── product
│   ├── handlers
│   │   └── product_handler.go
│   ├── services
│   │   └── product_service.go
│   ├── repositories
│   │   └── product_repository.go
│   └── models
│       └── product.go
├── database
│   └── db.go
├── cache
│   └── cache.go
├── tests
│   └── product_test.go
├── Dockerfile
├── docker-compose.yml
├── .env.sample
```

## API Specification
### Create a new product
- **Endpoint**: `POST /products`
- **Request Body**:
```json
{
  "name": "Product Name",
  "price": 100.0,
  "description": "Product Description"
}
```
- **Response**:

Success:
```json
{
  "id": 1,
  "name": "Product Name",
  "price": 100.0,
  "description": "Product Description"
}

```
Duplicate product name: 
```json
{
  "error": "Product with the same name already exists"
}
```
Input validation error:
```json
{
  "error": "Invalid input data"
}

Error creating product:
```json
{
  "error": "Error creating product"
}
```

### Get a product by ID
- **Endpoint**: `GET /products/{id}`
- **Response**:
Success:
```json
{
  "id": 1,
  "name": "Product Name",
  "price": 100.0,
  "description": "Product Description"
}
```
Not found:
```json
{
  "error": "Product not found"
}
``` 

Error retrieving product:
```json
{
  "error": "Error retrieving product"
}
```


