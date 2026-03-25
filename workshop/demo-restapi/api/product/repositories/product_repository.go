package repositories

import (
	"context"
	"database/sql"
	"errors"

	"api/product/models"
)

var (
	ErrDuplicateName = errors.New("product with the same name already exists")
	ErrNotFound      = errors.New("product not found")
)

type ProductRepository interface {
	Create(ctx context.Context, p *models.Product) error
	GetByID(ctx context.Context, id int) (*models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, p *models.Product) error {
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO products (name, price, description) VALUES ($1, $2, $3) RETURNING id",
		p.Name, p.Price, p.Description,
	).Scan(&p.ID)
	if err != nil {
		if isDuplicateKeyError(err) {
			return ErrDuplicateName
		}
		return err
	}
	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	p := &models.Product{}
	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, price, description FROM products WHERE id = $1", id,
	).Scan(&p.ID, &p.Name, &p.Price, &p.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return p, nil
}

func isDuplicateKeyError(err error) bool {
	return err != nil && contains(err.Error(), "duplicate key")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
