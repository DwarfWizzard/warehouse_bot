package service

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type ShopingCartService struct {
	repo repository.ShopingCartRepo
}

func NewShopingCartService(repo repository.ShopingCartRepo) *ShopingCartService {
	return &ShopingCartService{
		repo: repo,
	}
}

func (s *ShopingCartService) CreateCart(orderId int, productId int) error {
	return s.repo.Create(orderId, productId)
}

func (s *ShopingCartService) GetProductsFromCart(orderId int) ([]models.Product, error) {
	var products []models.Product

	products, err := s.repo.GetProducts(orderId)
	return products, err
}