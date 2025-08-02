package repositories

import (
	"go-linq-api/internal/helpers"
	"go-linq-api/internal/linq"
	"go-linq-api/internal/models"

	"gorm.io/gorm"
)

type WardRepository interface {
	GetAll() helpers.OperationResult
	GetWardDetails(pagination helpers.PaginationParam) helpers.OperationResult
}

type wardRepository struct {
	db *gorm.DB
}

func NewWardRepository(db *gorm.DB) WardRepository {
	return &wardRepository{db: db}
}

// ---------------- GET ALL ----------------
func (r *wardRepository) GetAll() helpers.OperationResult {
	var wards []models.Ward
	if err := r.db.Find(&wards).Error; err != nil {
		return helpers.NewOperationResultError(err.Error())
	}
	return helpers.NewOperationResultSuccess(wards)
}

// ---------------- GET WARD DETAILS WITH PAGINATION ----------------
func (r *wardRepository) GetWardDetails(pagination helpers.PaginationParam) helpers.OperationResult {
	pagination.Normalize()

	// Query cơ bản
	q := linq.From(r.db, &models.Ward{}).
		Select(`
			wards.code AS ward_code,
			wards.name AS ward_name,
			provinces.name AS province_name,
			administrative_units.full_name AS unit_name
		`).
		LeftJoin("provinces", "provinces.code = wards.province_code").
		LeftJoin("administrative_units", "administrative_units.id = wards.administrative_unit_id").
		OrderBy("province_name ASC, ward_name ASC")

	// Đếm tổng số bản ghi
	var total int64
	if err := q.Build().Count(&total).Error; err != nil {
		return helpers.NewOperationResultError(err.Error())
	}

	// Lấy dữ liệu theo trang
	var results []map[string]interface{}
	skip := (pagination.PageNumber - 1) * pagination.PageSize

	if err := q.Build().
		Offset(skip).
		Limit(pagination.PageSize).
		Scan(&results).Error; err != nil {
		return helpers.NewOperationResultError(err.Error())
	}

	// Đóng gói phân trang
	paged := helpers.CreatePaginationResult(results, int(total), pagination.PageNumber, pagination.PageSize, skip)
	return helpers.NewOperationResultSuccess(paged)
}
