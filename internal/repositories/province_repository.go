package repositories

import (
	"go-linq-api/internal/linq"
	"go-linq-api/internal/models"

	"gorm.io/gorm"
)

type ProvinceRepository interface {
	GetAll() ([]models.Province, error)
	GetByCode(code string) (*models.Province, error)
	GetWithJoins() ([]map[string]interface{}, error)
}

type provinceRepository struct {
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &provinceRepository{db: db}
}

func (r *provinceRepository) GetAll() ([]models.Province, error) {
	var provinces []models.Province
	err := r.db.Find(&provinces).Error
	return provinces, err
}

func (r *provinceRepository) GetByCode(code string) (*models.Province, error) {
	var province models.Province
	err := r.db.Where("code = ?", code).First(&province).Error
	if err != nil {
		return nil, err
	}
	return &province, nil
}

func (r *provinceRepository) GetWithJoins() ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	// Join provinces + administrative_units + wards
	err := linq.From(r.db, &models.Province{}).
		Select("provinces.code as province_code, provinces.name as province_name, administrative_units.full_name as unit_name, COUNT(wards.code) as ward_count").
		LeftJoin("administrative_units ON administrative_units.id = provinces.administrative_unit_id", "").
		LeftJoin("wards ON wards.province_code = provinces.code", "").
		GroupBy("provinces.code, provinces.name, administrative_units.full_name").
		ToList(&results)

	return results, err
}
