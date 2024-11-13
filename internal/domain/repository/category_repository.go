package repository

import (
	"catalog_service/internal/domain/model"
	"catalog_service/internal/domain/payload/params"
	"catalog_service/internal/domain/payload/request"
	"catalog_service/internal/domain/payload/response"
	"context"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	SaveCategory(ctx context.Context, request request.CreateCategoryRequest) (response.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (model.Category, error)
	GetAllCategories(ctx context.Context, params params.CategoryQueryParams) ([]model.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
	UpdateCategory(ctx context.Context, category model.Category) (response.CategoryResponse, error)
}
