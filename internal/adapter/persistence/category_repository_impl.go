package persistence

import (
	"catalog_service/internal/domain/model"
	"catalog_service/internal/domain/payload/params"
	"catalog_service/internal/domain/payload/request"
	response2 "catalog_service/internal/domain/payload/response"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type CategoryRepositoryInterface struct {
	db *gorm.DB
}

// NewCategoryRepositoryInterface creates a new instance of CategoryRepositoryInterface
func NewCategoryRepositoryInterface(db *gorm.DB) *CategoryRepositoryInterface {
	return &CategoryRepositoryInterface{db: db}
}

// SaveCategory inserts a new category into the database
func (repo *CategoryRepositoryInterface) SaveCategory(ctx context.Context, request request.CreateCategoryRequest) (response2.CategoryResponse, error) {
	category := model.Category{
		CategoryID:   uuid.New(),
		Name:         request.Name,
		Description:  request.Description,
		RestaurantID: request.RestaurantID,
		CountryCode:  request.CountryCode,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save to the database
	if err := repo.db.WithContext(ctx).Create(&category).Error; err != nil {
		return response2.CategoryResponse{}, err
	}

	// Convert the saved Category model to CategoryResponse
	categoryResponse := response2.CategoryResponse{
		StatusSchema: response2.StatusSchema{
			StatusCode: "SWT-00-000",
			StatusMessage: response2.StatusMessage{
				English:   "Success",
				Indonesia: "Berhasil",
			},
		},
		DataSchema: response2.DataSchema{
			Data: response2.CategoryData{
				CategoryID:   category.CategoryID,
				Name:         category.Name,
				Description:  category.Description,
				RestaurantID: category.RestaurantID,
				CountryCode:  category.CountryCode,
				CreatedAt:    category.CreatedAt,
				UpdatedAt:    category.UpdatedAt,
			},
			Pagination: response2.PaginationSchema{}, // No pagination in a single item response
		},
	}

	return categoryResponse, nil
}

// GetCategoryByID fetches a category by its ID
func (repo *CategoryRepositoryInterface) GetCategoryByID(ctx context.Context, id uuid.UUID) (model.Category, error) {
	var category model.Category
	if err := repo.db.WithContext(ctx).First(&category, "category_id = ?", id).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// GetAllCategories retrieves categories with optional filtering, pagination, sorting, and search
func (repo *CategoryRepositoryInterface) GetAllCategories(ctx context.Context, params params.CategoryQueryParams) ([]model.Category, error) {
	var categories []model.Category
	query := repo.db.WithContext(ctx).Model(&model.Category{})

	// Apply filters
	if params.RestaurantID != uuid.Nil {
		query = query.Where("restaurant_id = ?", params.RestaurantID)
	}
	if params.CountryCode != "" {
		query = query.Where("country_code = ?", params.CountryCode)
	}

	// Apply search
	if params.Search != "" {
		searchPattern := "%" + strings.ToLower(params.Search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", searchPattern, searchPattern)
	}

	// Apply sorting
	if params.SortBy == "" {
		params.SortBy = "created_at" // Default sorting field
	}
	if params.SortOrder != "asc" && params.SortOrder != "desc" {
		params.SortOrder = "asc" // Default sorting order
	}
	query = query.Order(params.SortBy + " " + params.SortOrder)

	// Apply pagination
	if params.Page > 0 && params.PageSize > 0 {
		offset := (params.Page - 1) * params.PageSize
		query = query.Offset(offset).Limit(params.PageSize)
	}

	// Execute query
	if err := query.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// UpdateCategory updates an existing category
func (repo *CategoryRepositoryInterface) UpdateCategory(ctx context.Context, category model.Category) (model.Category, error) {
	if err := repo.db.WithContext(ctx).Save(&category).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// DeleteCategory removes a category by its ID
func (repo *CategoryRepositoryInterface) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return repo.db.WithContext(ctx).Delete(&model.Category{}, "category_id = ?", id).Error
}
