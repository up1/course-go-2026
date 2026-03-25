package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"api/product/handlers"
	"api/product/models"
	"api/product/repositories"
	"api/product/services"

	"github.com/gin-gonic/gin"
)

// --- Fake cache (always misses) ---
type fakeCache struct{}

func (f *fakeCache) Get(context.Context, string, any) error                { return fmt.Errorf("miss") }
func (f *fakeCache) Set(context.Context, string, any, time.Duration) error { return nil }
func (f *fakeCache) Delete(context.Context, string) error                  { return nil }

// --- Fake repository ---

type fakeRepo struct {
	products map[int]*models.Product
	nextID   int
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{products: make(map[int]*models.Product), nextID: 1}
}

func (r *fakeRepo) Create(_ context.Context, p *models.Product) error {
	for _, existing := range r.products {
		if existing.Name == p.Name {
			return repositories.ErrDuplicateName
		}
	}
	p.ID = r.nextID
	stored := *p
	r.products[r.nextID] = &stored
	r.nextID++
	return nil
}

func (r *fakeRepo) GetByID(_ context.Context, id int) (*models.Product, error) {
	p, ok := r.products[id]
	if !ok {
		return nil, repositories.ErrNotFound
	}
	return p, nil
}

// --- Helpers ---

func setupRouter(repo repositories.ProductRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	svc := services.NewProductService(repo, &fakeCache{})
	h := handlers.NewProductHandler(svc)
	r := gin.New()
	r.POST("/products", h.Create)
	r.GET("/products/:id", h.GetByID)
	return r
}

// --- Tests ---

func TestCreateProduct_Success(t *testing.T) {
	router := setupRouter(newFakeRepo())

	body, _ := json.Marshal(map[string]any{
		"name": "Laptop", "price": 999.99, "description": "A laptop",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", w.Code, w.Body.String())
	}

	var p models.Product
	json.Unmarshal(w.Body.Bytes(), &p)
	if p.ID == 0 || p.Name != "Laptop" {
		t.Fatalf("unexpected product: %+v", p)
	}
}

func TestCreateProduct_InvalidInput(t *testing.T) {
	router := setupRouter(newFakeRepo())

	body, _ := json.Marshal(map[string]any{"name": ""})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestCreateProduct_Duplicate(t *testing.T) {
	repo := newFakeRepo()
	router := setupRouter(repo)

	body, _ := json.Marshal(map[string]any{
		"name": "Laptop", "price": 999.99, "description": "A laptop",
	})

	// First create
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Duplicate
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("expected 409, got %d", w.Code)
	}
}

func TestGetProduct_Success(t *testing.T) {
	repo := newFakeRepo()
	repo.Create(context.Background(), &models.Product{
		Name: "Mouse", Price: 29.99, Description: "A mouse",
	})
	router := setupRouter(repo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var p models.Product
	json.Unmarshal(w.Body.Bytes(), &p)
	if p.Name != "Mouse" {
		t.Fatalf("unexpected product: %+v", p)
	}
}

func TestGetProduct_NotFound(t *testing.T) {
	router := setupRouter(newFakeRepo())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestGetProduct_InvalidID(t *testing.T) {
	router := setupRouter(newFakeRepo())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products/abc", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
