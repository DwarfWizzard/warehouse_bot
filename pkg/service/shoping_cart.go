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

func (s *ShopingCartService) CreateCart(orderId int, productId int, productPrice int, deliveryFormat string) error {
	var cart models.ShopingCart
	cart, err := s.repo.GetCart(orderId, productId)
	if err != nil && (err.Error() == "sql: no rows in result set" || cart.OrderId==0) {
		return s.repo.Create(orderId, productId, productPrice, deliveryFormat)
	} 
	return err
}

func (s *ShopingCartService) GetProductsFromCart(orderId int) ([]models.Product, error) {
	var products []models.Product

	products, err := s.repo.GetProductsFromCart(orderId)
	if err != nil {
		return products, err
	}

	return products, err
}

func (s *ShopingCartService) GetQuantity(orderId int, productId int) (int, error) {
	return s.repo.GetQuantity(orderId, productId)
}

func (s *ShopingCartService) UpdateQuantity(orderId int, productId int, quantity int) error {
	return s.repo.UpdateQuantity(orderId, productId, quantity)
}

func (s *ShopingCartService) DeleteCart(orderId int) error {
	return s.repo.DeleteCart(orderId)
}

func (s *ShopingCartService) DeleteProductFromCart(orderId int, productId int) error {
	return s.repo.DeleteProductFromCart(orderId, productId)
}

func (s *ShopingCartService) GetCart(orderId int, productId int) (models.ShopingCart, error) {
	return s.repo.GetCart(orderId, productId)
}

func (s *ShopingCartService) GetCarts(orderId int) ([]models.ShopingCart, error) {
	return s.repo.GetCarts(orderId)
}