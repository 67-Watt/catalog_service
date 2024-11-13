package response

import (
	"github.com/google/uuid"
	"time"
)

// CategoryData represents the structure for category-specific data
type CategoryData struct {
	CategoryID   uuid.UUID `json:"category_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	RestaurantID uuid.UUID `json:"restaurant_id"`
	CountryCode  string    `json:"country_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CategoryResponse represents the full response for a paginated list of categories
type CategoryResponse struct {
	StatusSchema StatusSchema `json:"status_schema"`
	DataSchema   DataSchema   `json:"data_schema"`
}

// NewCategoryResponse creates a new paginated response for categories
func NewCategoryResponse(categories []CategoryData, totalCount, currentPage, pageSize int) CategoryResponse {
	totalPages := (totalCount + pageSize - 1) / pageSize // Calculate total pages

	return CategoryResponse{
		StatusSchema: StatusSchema{
			StatusCode: "SWT-00-000",
			StatusMessage: StatusMessage{
				English:   "Success",
				Indonesia: "Berhasil",
			},
		},
		DataSchema: DataSchema{
			Data: categories,
			Pagination: PaginationSchema{
				TotalCount:  totalCount,
				CurrentPage: currentPage,
				TotalPages:  totalPages,
				PageSize:    pageSize,
			},
		},
	}
}
