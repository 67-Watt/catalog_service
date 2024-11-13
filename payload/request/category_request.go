package request

import (
	"github.com/google/uuid"
)

// CreateCategoryRequest represents the request payload for creating a new category
type CreateCategoryRequest struct {
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description,omitempty"`
	RestaurantID uuid.UUID `json:"restaurant_id" binding:"required"`
	CountryCode  string    `json:"country_code" binding:"required,len=2"`
}

// UpdateCategoryRequest represents the request payload for updating an existing category
type UpdateCategoryRequest struct {
	CategoryID   uuid.UUID `json:"category_id" binding:"required"` // Used to identify which category to update
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	RestaurantID uuid.UUID `json:"restaurant_id,omitempty"`
	CountryCode  string    `json:"country_code,omitempty"`
}

// CategoryFilterRequest represents the query parameters for filtering, sorting, and paginating categories
//type CategoryFilterRequest struct {
//	RestaurantID uuid.UUID `form:"restaurant_id"`
//	CountryCode  string    `form:"country_code"`
//	Search       string    `form:"search"`
//	Page         int       `form:"page" binding:"required,min=1"`
//	PageSize     int       `form:"page_size" binding:"required,min=1"`
//	SortBy       string    `form:"sort_by"`
//	SortOrder    string    `form:"sort_order" binding:"oneof=asc desc"`
//}
