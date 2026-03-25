package services

import (
	"context"
	"errors"
	"time"

	"api/cache"
	"api/product/models"
	"api/product/repositories"

	"github.com/redis/go-redis/v9"
)

const cacheTTL = 10 * time.Minute

type ProductService interface {
	Create(ctx context.Context, p *models.Product) error
	GetByID(ctx context.Context, id int) (*models.Product, error)
}

type productService struct {
	repo  repositories.ProductRepository
	cache cache.Cache
}

func NewProductService(repo repositories.ProductRepository, cache cache.Cache) ProductService {
	return &productService{repo: repo, cache: cache}
}

func (s *productService) Create(ctx context.Context, p *models.Product) error {
	return s.repo.Create(ctx, p)
}

func (s *productService) GetByID(ctx context.Context, id int) (*models.Product, error) {
	var p models.Product

	// Try to get from cache first
	key := cache.ProductKey(id)

	err := s.cache.Get(ctx, key, &p)
	if err == nil {
		return &p, nil
	}

	// If cache miss, get from database
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Set to cache (ignore cache errors)
	if cacheErr := s.cache.Set(ctx, key, product, cacheTTL); cacheErr != nil {
		if !errors.Is(cacheErr, redis.Nil) {
			// log but don't fail the request
		}
	}

	return product, nil
}
