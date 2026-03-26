package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api/database"
	"api/product/handlers"
	"api/product/repositories"
	"api/product/services"

	"github.com/gin-gonic/gin"
)

func setupRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db, err := database.NewDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)
	handler := handlers.NewProductHandler(service)

	r := gin.Default()
	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProduct)

	return r
}

func TestCreateProduct(t *testing.T) {
	router := setupRouter(t)

	body := map[string]any{
		"name":        "Test Product",
		"price":       29.99,
		"description": "A test product",
	}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["name"] != "Test Product" {
		t.Errorf("expected name 'Test Product', got '%v'", response["name"])
	}
	if response["id"] == nil || response["id"].(float64) == 0 {
		t.Error("expected a valid product ID")
	}
}

func TestCreateProductValidationError(t *testing.T) {
	router := setupRouter(t)

	body := map[string]any{"name": ""}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["error"] != "Invalid input data" {
		t.Errorf("expected 'Invalid input data', got '%v'", response["error"])
	}
}

func TestCreateDuplicateProduct(t *testing.T) {
	router := setupRouter(t)

	body := map[string]any{
		"name":        "Duplicate",
		"price":       10.0,
		"description": "First",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	req, _ = http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("expected status %d, got %d", http.StatusConflict, w.Code)
	}

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["error"] != "Product with the same name already exists" {
		t.Errorf("expected duplicate error message, got '%v'", response["error"])
	}
}

func TestGetProduct(t *testing.T) {
	router := setupRouter(t)

	body := map[string]any{
		"name":        "Fetchable",
		"price":       50.0,
		"description": "To be fetched",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/products/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["name"] != "Fetchable" {
		t.Errorf("expected name 'Fetchable', got '%v'", response["name"])
	}
}

func TestGetProductNotFound(t *testing.T) {
	router := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}

	var response map[string]any
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["error"] != "Product not found" {
		t.Errorf("expected 'Product not found', got '%v'", response["error"])
	}
}
