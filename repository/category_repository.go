package repository

import (
	"catalog_service/models"
	"catalog_service/payload/params"
	"catalog_service/payload/request"
	"catalog_service/payload/response"
	"context"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	SaveCategory(ctx context.Context, request request.CreateCategoryRequest) (response.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (models.Category, error)
	GetAllCategories(ctx context.Context, params params.CategoryQueryParams) ([]models.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
	UpdateCategory(ctx context.Context, category models.Category) (response.CategoryResponse, error)
}
