package pagination

import (
	"net/http"
	"strconv"
)

var (
	// DefaultPageSize specifies the default page size
	DefaultPageSize = 100
	// MaxPageSize specifies the maximum page size
	MaxPageSize = 5000
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PageSizeVar specifies the query parameter name for page size
	PageSizeVar = "size"

	Sort = "sort"

	SortBy = "sortBy"
)

// Pages represents a paginated list of data items.
type Pages struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"size"`
	PageCount  int         `json:"page_count"`
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
	Sort       string      `json:sort`
	SortBy     string      `json:sortBy`
}

// New creates a new Pages instance.
// The page parameter is 1-based and refers to the current page index/number.
// The perPage parameter refers to the number of items on each page.
// And the total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func New(page, perPage, total int, sort, sortBy string) *Pages {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}
	if len(sort) > 0 && len(sortBy) == 0 {
		sortBy = "asc"
	}
	return &Pages{
		Page:       page,
		PerPage:    perPage,
		PageCount:  pageCount,
		TotalCount: total,
		Sort:       sort,
		SortBy:     sortBy,
	}
}

// NewFromRequest creates a Pages object using the query parameters found in the given HTTP request.
// count stands for the total number of items. Use -1 if this is unknown.
func NewFromRequest(req *http.Request, count int) *Pages {
	page := parseInt(req.URL.Query().Get(PageVar), 1)
	perPage := parseInt(req.URL.Query().Get(PageSizeVar), DefaultPageSize)

	sort := req.URL.Query().Get(Sort)
	sortBy := req.URL.Query().Get(SortBy)
	return New(page, perPage, count, sort, sortBy)
}

// parseInt parses a string into an integer. If parsing is failed, defaultValue will be returned.
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// Offset returns the OFFSET value that can be used in a SQL statement.
func (p *Pages) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// Limit returns the LIMIT value that can be used in a SQL statement.
func (p *Pages) Limit() int {
	return p.PerPage
}
