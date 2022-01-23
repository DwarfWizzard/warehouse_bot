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
	if err != nil {
		return products, err
	}

	for i := range products {
		products[i].Quantity, err = s.repo.GetQuantity(orderId, products[i].Id)
		if err != nil {
			return products, err
		}
	}

	return products, err
}

func (s *ShopingCartService) GetQuantity(orderId int, productId int) (int, error) {
	return s.repo.GetQuantity(orderId, productId)
}

func (s *ShopingCartService) UpdateQuantity(orderId int, productId int, quantity int) error {
	return s.repo.UpdateQuantity(orderId, productId, quantity)
}