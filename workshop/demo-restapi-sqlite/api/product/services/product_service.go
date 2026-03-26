package services

import (
	"database/sql"
	"errors"

	"api/product/models"
	"api/product/repositories"
)

var (
	ErrProductNotFound = errors.New("Product not found")
	ErrDuplicateName   = errors.New("Product with the same name already exists")
	ErrCreating        = errors.New("Error creating product")
	ErrRetrieving      = errors.New("Error retrieving product")
)

type ProductService interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetProductByID(id int64) (models.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	created, err := s.repo.Create(product)
	if err != nil {
		if errors.Is(err, repositories.ErrDuplicateProduct) {
			return models.Product{}, ErrDuplicateName
		}
		return models.Product{}, ErrCreating
	}
	return created, nil
}

func (s *productService) GetProductByID(id int64) (models.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Product{}, ErrProductNotFound
		}
		return models.Product{}, ErrRetrieving
	}
	return product, nil
}
