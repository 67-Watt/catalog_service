package params

import "github.com/google/uuid"

// CategoryQueryParams holds filtering, pagination, and sorting parameters for category queries
type CategoryQueryParams struct {
	RestaurantID uuid.UUID `json:"restaurant_id"` // Filter by restaurant ID
	CountryCode  string    `json:"country_code"`  // Filter by country code
	Search       string    `json:"search"`        // Search term for name or description
	Page         int       `json:"page"`          // Pagination: current page number
	PageSize     int       `json:"page_size"`     // Pagination: number of items per page
	SortBy       string    `json:"sort_by"`       // Sorting field, e.g., "name" or "created_at"
	SortOrder    string    `json:"sort_order"`    // Sorting order, either "asc" or "desc"
}
