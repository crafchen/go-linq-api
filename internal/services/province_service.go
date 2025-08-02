package services

import (
	"go-linq-api/internal/models"
	"go-linq-api/internal/repositories"
)

type ProvinceService interface {
	GetAll() ([]models.Province, error)
	GetByCode(code string) (*models.Province, error)
	GetWithStatistics() ([]map[string]interface{}, error)
}

type provinceService struct {
	repo repositories.ProvinceRepository
}

func NewProvinceService(repo repositories.ProvinceRepository) ProvinceService {
	return &provinceService{repo: repo}
}

func (s *provinceService) GetAll() ([]models.Province, error) {
	return s.repo.GetAll()
}

func (s *provinceService) GetByCode(code string) (*models.Province, error) {
	return s.repo.GetByCode(code)
}

func (s *provinceService) GetWithStatistics() ([]map[string]interface{}, error) {
	return s.repo.GetWithJoins()
}
