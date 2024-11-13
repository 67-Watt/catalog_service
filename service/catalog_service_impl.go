package service

import (
	"catalog_service/payload/params"
	"catalog_service/payload/request"
	"catalog_service/payload/response"
	"catalog_service/repository"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

// categoryService is the implementation of CategoryService
type categoryService struct {
	repo repository.CategoryRepository
}

// NewCategoryService creates a new instance of categoryService
func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

// CreateCategory handles the creation of a new category and returns a CategoryResponse
func (s *categoryService) CreateCategory(ctx context.Context, input request.CreateCategoryRequest) (response.CategoryResponse, error) {
	if input.Name == "" {
		return response.CategoryResponse{}, errors.New("category name is required")
	}

	// Use repository to save and get the CategoryResponse
	categoryResponse, err := s.repo.SaveCategory(ctx, input)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	return categoryResponse, nil
}

// GetCategoryByID retrieves a category by its ID and returns a CategoryResponse
func (s *categoryService) GetCategoryByID(ctx context.Context, id uuid.UUID) (response.CategoryResponse, error) {
	category, err := s.repo.GetCategoryByID(ctx, id)
	if err != nil {
		return response.CategoryResponse{}, err
	}

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
		},
	}

	return categoryResponse, nil
}

// ListCategories retrieves categories based on the provided filter and returns a paginated response
func (s *categoryService) ListCategories(ctx context.Context, param params.CategoryQueryParams) (response.CategoryResponse, error) {
	categories, err := s.repo.GetAllCategories(ctx, param)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	totalCount := len(categories)

	categoryData := make([]response.CategoryData, len(categories))
	for i, category := range categories {
		categoryData[i] = response.CategoryData{
			CategoryID:   category.CategoryID,
			Name:         category.Name,
			Description:  category.Description,
			RestaurantID: category.RestaurantID,
			CountryCode:  category.CountryCode,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
		}
	}

	categoryResponse := response.CategoryResponse{
		StatusSchema: response.StatusSchema{
			StatusCode: "SWT-00-000",
			StatusMessage: response.StatusMessage{
				English:   "Success",
				Indonesia: "Berhasil",
			},
		},
		DataSchema: response.DataSchema{
			Data: categoryData,
			Pagination: response.PaginationSchema{
				TotalCount:  totalCount,
				CurrentPage: param.Page,
				TotalPages:  (totalCount + param.PageSize - 1) / param.PageSize,
				PageSize:    param.PageSize,
			},
		},
	}

	return categoryResponse, nil
}

// UpdateCategory updates an existing category and returns a CategoryResponse
func (s *categoryService) UpdateCategory(ctx context.Context, id uuid.UUID, input request.UpdateCategoryRequest) (response.CategoryResponse, error) {
	category, err := s.repo.GetCategoryByID(ctx, id)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	if input.Name != "" {
		category.Name = input.Name
	}
	if input.Description != "" {
		category.Description = input.Description
	}
	if input.RestaurantID != uuid.Nil {
		category.RestaurantID = input.RestaurantID
	}
	if input.CountryCode != "" {
		category.CountryCode = input.CountryCode
	}
	category.UpdatedAt = time.Now()

	// Update the category in the repository
	updatedCategoryResponse, err := s.repo.UpdateCategory(ctx, category)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	return updatedCategoryResponse, nil
}

// DeleteCategory deletes a category by its ID
func (s *categoryService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCategory(ctx, id)
}
