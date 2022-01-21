package service

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type ProductsService struct {
	repo repository.ProductsRepo
}

func NewProductsService(repo repository.ProductsRepo) *ProductsService {
	return &ProductsService{
		repo: repo,
	}
}

func (s *ProductsService) GetProducts(offset int) ([]models.Product, error) {
	return s.repo.GetProducts(offset)
}

func (s *ProductsService) CountProducts() (int, error) {
	return s.repo.CountProducts()
}

func (s *ProductsService) CountProductsOnPage(offset int) (int, error) {
	return s.repo.CountProductsOnPage(offset)
}
