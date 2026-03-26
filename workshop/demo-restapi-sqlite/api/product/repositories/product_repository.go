package repositories

import (
	"database/sql"
	"errors"

	"api/product/models"
)

var ErrDuplicateProduct = errors.New("product with the same name already exists")

type ProductRepository interface {
	Create(product models.Product) (models.Product, error)
	GetByID(id int64) (models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	result, err := r.db.Exec(
		"INSERT INTO products (name, price, description) VALUES (?, ?, ?)",
		product.Name, product.Price, product.Description,
	)
	if err != nil {
		if isUniqueConstraintError(err) {
			return models.Product{}, ErrDuplicateProduct
		}
		return models.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}

	product.ID = id
	return product, nil
}

func (r *productRepository) GetByID(id int64) (models.Product, error) {
	var product models.Product
	err := r.db.QueryRow(
		"SELECT id, name, price, description FROM products WHERE id = ?", id,
	).Scan(&product.ID, &product.Name, &product.Price, &product.Description)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Product{}, sql.ErrNoRows
		}
		return models.Product{}, err
	}

	return product, nil
}

func isUniqueConstraintError(err error) bool {
	return err.Error() == "UNIQUE constraint failed: products.name"
}
