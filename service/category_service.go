package service

import (
	"catalog_service/payload/request"
	"catalog_service/payload/response"
	"context"
	"github.com/google/uuid"
)

// CategoryService defines the methods for category-related business operations
type CategoryService interface {
	CreateCategory(ctx context.Context, input request.CreateCategoryRequest) (response.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (response.CategoryResponse, error)
	ListCategories(ctx context.Context, filter request.CategoryFilterRequest) (response.CategoryResponse, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, input request.UpdateCategoryRequest) (response.CategoryResponse, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}
