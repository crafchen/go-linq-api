package services

import (
	"go-linq-api/internal/models"
	"go-linq-api/internal/repositories"
)

type WardService interface {
	GetAll() ([]models.Ward, error)
	GetWardDetails() ([]map[string]interface{}, error)
}

type wardService struct {
	repo repositories.WardRepository
}

func NewWardService(repo repositories.WardRepository) WardService {
	return &wardService{repo: repo}
}

func (s *wardService) GetAll() ([]models.Ward, error) {
	return s.repo.GetAll()
}

func (s *wardService) GetWardDetails() ([]map[string]interface{}, error) {
	return s.repo.GetWardDetails()
}
