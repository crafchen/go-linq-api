package helpers

import (
	"math"
)

// PaginationResult trả về thông tin phân trang
type PaginationResult struct {
	TotalCount int         `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
	PageNumber int         `json:"pageNumber"`
	PageSize   int         `json:"pageSize"`
	Skip       int         `json:"skip"`
	Result     interface{} `json:"result"`
}

// CreatePaginationResult tạo dữ liệu phân trang
func CreatePaginationResult(items interface{}, count int, pageNumber, pageSize, skip int) PaginationResult {
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	return PaginationResult{
		TotalCount: count,
		TotalPage:  totalPage,
		PageNumber: pageNumber,
		PageSize:   pageSize,
		Skip:       skip,
		Result:     items,
	}
}

// PaginationUtility chứa kết quả + meta phân trang
type PaginationUtility[T any] struct {
	Pagination PaginationResult `json:"pagination"`
	Result     []T              `json:"result"`
}

// NewPaginationUtility khởi tạo đối tượng phân trang từ dữ liệu đã có
func NewPaginationUtility[T any](items []T, count int, pageNumber int, pageSize int, skip int) PaginationUtility[T] {
	return PaginationUtility[T]{
		Result:     items,
		Pagination: CreatePaginationResult(items, count, pageNumber, pageSize, skip),
	}
}

// Create tạo PaginationUtility từ slice với phân trang cục bộ
func Create[T any](source []T, pageNumber int, pageSize int, isPaging bool) PaginationUtility[T] {
	count := len(source)
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	skip := (pageNumber - 1) * pageSize
	var items []T
	if isPaging {
		end := skip + pageSize
		if skip > count {
			items = []T{}
		} else if end > count {
			items = source[skip:]
		} else {
			items = source[skip:end]
		}
	} else {
		items = source
	}

	return NewPaginationUtility(items, count, pageNumber, pageSize, skip)
}

// ----------------- PaginationParam -----------------

// PaginationParam là tham số phân trang
type PaginationParam struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
}

// MaxPageSize giới hạn số record trên 1 trang
const MaxPageSize = 50

// NewPaginationParam khởi tạo với mặc định
func NewPaginationParam(pageNumber int, pageSize int) PaginationParam {
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	return PaginationParam{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}
}

// Normalize tự động điều chỉnh giá trị về chuẩn
func (p *PaginationParam) Normalize() {
	if p.PageNumber <= 0 {
		p.PageNumber = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > MaxPageSize {
		p.PageSize = MaxPageSize
	}
}
