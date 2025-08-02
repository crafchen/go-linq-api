package services

import (
	"go-linq-api/internal/helpers"
	"go-linq-api/internal/repositories"
)

type WardService interface {
	GetAll() helpers.OperationResult
	GetWardDetails(pagination helpers.PaginationParam) helpers.OperationResult
}

type wardService struct {
	repo repositories.WardRepository
}

func NewWardService(repo repositories.WardRepository) WardService {
	return &wardService{repo: repo}
}

// Lấy tất cả wards
func (s *wardService) GetAll() helpers.OperationResult {
	wards := s.repo.GetAll()
	return helpers.NewOperationResultSuccess(wards)
}

// Lấy ward details có phân trang
func (s *wardService) GetWardDetails(pagination helpers.PaginationParam) helpers.OperationResult {
	return s.repo.GetWardDetails(pagination)
}
