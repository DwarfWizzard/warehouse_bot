package service

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type Users interface {
	CreateUser(telegramId int64) error
	GetUser(telegramId int64) (models.User, error)
	UpdateUser(telegramId int64, field string, value string) error
	UpdateUserStatus(telegramId int64, status string) error
}

type Products interface {
	GetProducts(offset int) ([]models.Product, error)
	CountProducts() (int, error)
	CountProductsOnPage(offset int) (int, error)
}

type ShopingCart interface {
	CreateCart(orderId int, productId int) error 
	GetProductsFromCart(orderId int) ([]models.Product, error)
}

type Order interface{
	GetOrder(telegramId int64) (models.Order, error)
	UpdateOrder(telegramId int64, field string, value string) error
}

type Service struct {
	Users
	Products
	ShopingCart
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users: NewUserService(repos.UsersRepo),
		Products: NewProductsService(repos.ProductsRepo),
		ShopingCart: NewShopingCartService(repos.ShopingCartRepo),
		Order: NewOrderService(repos.OrderRepo),
	}
}