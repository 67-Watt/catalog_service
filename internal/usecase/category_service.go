package usecase

import (
	request2 "catalog_service/internal/domain/payload/request"
	"catalog_service/internal/domain/payload/response"
	"catalog_service/payload/request"
	"context"
	"github.com/google/uuid"
)

// CategoryService defines the methods for category-related business operations
type CategoryService interface {
	CreateCategory(ctx context.Context, input request2.CreateCategoryRequest) (response.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (response.CategoryResponse, error)
	ListCategories(ctx context.Context, filter request.CategoryFilterRequest) (response.CategoryResponse, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, input request2.UpdateCategoryRequest) (response.CategoryResponse, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}
