package repositories

import (
	"go-linq-api/internal/linq"
	"go-linq-api/internal/models"

	"gorm.io/gorm"
)

type WardRepository interface {
	GetAll() ([]models.Ward, error)
	GetWardDetails() ([]map[string]interface{}, error)
}

type wardRepository struct {
	db *gorm.DB
}

func NewWardRepository(db *gorm.DB) WardRepository {
	return &wardRepository{db: db}
}

func (r *wardRepository) GetAll() ([]models.Ward, error) {
	var wards []models.Ward
	err := r.db.Find(&wards).Error
	return wards, err
}

// Join: wards -> provinces -> administrative_units
func (r *wardRepository) GetWardDetails() ([]map[string]interface{}, error) {
	var results []map[string]interface{}

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

	// Dùng Scan để lấy map[string]interface{}
	err := q.Build().Scan(&results).Error
	return results, err
}
