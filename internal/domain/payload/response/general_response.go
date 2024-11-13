package response

type StatusMessage struct {
	English   string `json:"english"`
	Indonesia string `json:"indonesia"`
}

type StatusSchema struct {
	StatusCode    string        `json:"status_code"`
	StatusMessage StatusMessage `json:"status_message"`
}

type PaginationSchema struct {
	TotalCount  int `json:"total_count"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	PageSize    int `json:"page_size"`
}

type DataSchema struct {
	Data       interface{}      `json:"data"` // Can hold any list of data (e.g., categories, products)
	Pagination PaginationSchema `json:"pagination"`
}
