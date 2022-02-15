package repository

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type UsersRepo interface {
	Create(telegramId int64) error
	GetUser(telegramId int64) (models.User, error)
	UpdateUser(telegramId int64, field string, value string) error
	DeleteUser(telegramId int64) error
}

type CouriersRepo interface {
	Create(telegramId int64) error
	GetCourier(telegramId int64) (models.Courier, error)
	UpdateCourier(telegramId int64, field string, value string) error
}

type CouriersOrdersRepo interface {
	Create(orderId int, courierId int) error
	GetCourier(orderId int) (models.Courier, error)
	GetActiveOrders(courierId int) ([]models.Order, error)
	GetCourierOrders(courierId int) ([]models.Order, error)
	Update(orderId int, field string, value string) error
}

type ProductsRepo interface {
	GetProduct(productId int) (models.Product, error)
	GetProducts(offset int) ([]models.Product, error)
	CountProducts() (int, error)
	CountProductsOnPage(offset int) (int, error)
}

type ShopingCartRepo interface {
	Create(orderId int, productId int, productPrice int, deliveryFormat string) error
	GetCart(orderId int, productId int) (models.ShopingCart, error)
	GetCarts(orderId int) ([]models.ShopingCart, error)
	GetProductsFromCart(orderId int) ([]models.Product, error)
	GetQuantity(orderId int, productId int) (int, error)
	UpdateQuantity(orderId int, productId int, quantity int) error
	DeleteCart(orderId int) error
	DeleteProductFromCart(orderId int, productId int) error
}

type OrderRepo interface {
	Create(telegramId int64, date string) (models.Order, error)
	GetOrderById(orderId int) (models.Order, error)
	GetOrderByUser(telegramId int64) (models.Order, error)
	GetOrderUser(orderId int) (models.User, error)
	GetAllOrdersUser(userId int) ([]models.Order, error)
	UpdateOrder(telegramId int64, field string, value string) error
}

type SubsidiaryRepo interface{
	GetSubsidiary(cityName string) (models.Subsidiary, error)
	GetSubsidiarys() ([]models.Subsidiary, error)
}


type Repository struct {
	UsersRepo
	CouriersRepo
	CouriersOrdersRepo
	ProductsRepo
	ShopingCartRepo
	OrderRepo
	SubsidiaryRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UsersRepo:       NewUsersPostgres(db),
		CouriersRepo:    NewCourierPostgres(db),
		CouriersOrdersRepo: NewCouriersOrdersPostgres(db),
		ProductsRepo:    NewProductsPostgres(db),
		ShopingCartRepo: NewShopingCartPostgres(db),
		OrderRepo:       NewOrderPostgres(db),
		SubsidiaryRepo: NewSubsidiaryPostgres(db),
	}
}
