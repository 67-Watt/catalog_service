package repository

import (
	"catalog_service/models"
	"catalog_service/payload/params"
	"catalog_service/payload/request"
	"catalog_service/payload/response"
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
func (repo *CategoryRepositoryInterface) SaveCategory(ctx context.Context, request request.CreateCategoryRequest) (response.CategoryResponse, error) {
	category := models.Category{
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
		return response.CategoryResponse{}, err
	}

	// Convert the saved Category model to CategoryResponse
	categoryResponse := response.CategoryResponse{
		StatusSchema: response.StatusSchema{
			StatusCode: "SWT-00-000",
			StatusMessage: response.StatusMessage{
				English:   "Success",
				Indonesia: "Berhasil",
			},
		},
		DataSchema: response.DataSchema{
			Data: response.CategoryData{
				CategoryID:   category.CategoryID,
				Name:         category.Name,
				Description:  category.Description,
				RestaurantID: category.RestaurantID,
				CountryCode:  category.CountryCode,
				CreatedAt:    category.CreatedAt,
				UpdatedAt:    category.UpdatedAt,
			},
			Pagination: response.PaginationSchema{}, // No pagination in a single item response
		},
	}

	return categoryResponse, nil
}

// GetCategoryByID fetches a category by its ID
func (repo *CategoryRepositoryInterface) GetCategoryByID(ctx context.Context, id uuid.UUID) (models.Category, error) {
	var category models.Category
	if err := repo.db.WithContext(ctx).First(&category, "category_id = ?", id).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

// GetAllCategories retrieves categories with optional filtering, pagination, sorting, and search
func (repo *CategoryRepositoryInterface) GetAllCategories(ctx context.Context, params params.CategoryQueryParams) ([]models.Category, error) {
	var categories []models.Category
	query := repo.db.WithContext(ctx).Model(&models.Category{})

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
func (repo *CategoryRepositoryInterface) UpdateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	if err := repo.db.WithContext(ctx).Save(&category).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

// DeleteCategory removes a category by its ID
func (repo *CategoryRepositoryInterface) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return repo.db.WithContext(ctx).Delete(&models.Category{}, "category_id = ?", id).Error
}
