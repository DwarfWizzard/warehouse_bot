package repository

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type UsersRepo interface {
	Create(telegramId int64) error
	GetUser(telegramId int64) (models.User, error)
	UpdateUser(telegramId int64, field string, value string) error
}

type ProductsRepo interface {
	GetProduct(productId int) (models.Product, error)
	GetProducts(offset int) ([]models.Product, error)
	CountProducts() (int, error)
	CountProductsOnPage(offset int) (int, error) 
}

type ShopingCartRepo interface {
	Create(orderId int, productId int) error
	GetCart(orderId int, productId int) (models.ShopingCart, error)
	GetProductsFromCart(orderId int) ([]models.Product, error)
	GetQuantity(orderId int, productId int) (int, error)
	UpdateQuantity(orderId int, productId int, quantity int) error
	DeleteCart(orderId int) error
	DeleteProductFromCart(orderId int, productId int) error
}

type OrderRepo interface{
	Create(telegramId int64, date string) (models.Order, error)
	GetOrder(telegramId int64) (models.Order, error)
	UpdateOrder(telegramId int64, field string, value string) error
}

type Repository struct {
	UsersRepo
	ProductsRepo
	ShopingCartRepo
	OrderRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UsersRepo: NewUsersSQLite3(db),
		ProductsRepo: NewProductsSQLite3(db),
		ShopingCartRepo: NewShopingCartSQLite3(db),
		OrderRepo: NewOrderSQLite3(db),
	}
}
