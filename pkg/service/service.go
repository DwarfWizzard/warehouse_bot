package service

import (
	"os"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type Users interface {
	CreateUser(telegramId int64) error
	GetUser(telegramId int64) (models.User, error)
	UpdateUser(telegramId int64, field string, value string) error
	DeleteUser(telegramId int64) error
}

type Couriers interface {
	CreateCourier(telegramId int64) error
	GetCourier(telegramId int64) (models.Courier, error)
	UpdateCourier(telegramId int64, field string, value string) error
}

type CouriersOrders interface {
	CreateCourierOrder(orderId int, courierId int) error
	GetOrderCourier(orderId int) (models.Courier, error)
	GetActiveOrders(courierId int) ([]models.Order, error)
	GetOrders(courierId int) ([]models.Order, error)
	UpdateCourierOrder(orderId int, field string, value string) error
	GetOrderStatus(orderId int) (string, error)
}

type Products interface {
	GetProduct(productId int) (models.Product, error)
	GetProducts(offset int) ([]models.Product, error)
	CountProducts() (int, error)
	CountProductsOnPage(offset int) (int, error)
}

type ShopingCart interface {
	CreateCart(orderId int, productId int) error 
	GetProductsFromCart(orderId int) ([]models.Product, error)
	GetCart(orderId int, productId int) (models.ShopingCart, error)
	GetQuantity(orderId int, productId int) (int, error)
	UpdateQuantity(orderId int, productId int, quantity int) error
	DeleteCart(orderId int) error
	DeleteProductFromCart(orderId int, productId int) error
}

type Order interface{
	GetOrderById(orderId int) (models.Order, error)
	GetOrderByUser(telegramId int64) (models.Order, error)
	UpdateOrder(telegramId int64, field string, value string) error
	GetOrderUser(orderId int) (models.User, error)
}

type Logger interface{
	PrintLog(message string, flag int)
}

type Service struct {
	Users
	Couriers
	CouriersOrders
	Products
	ShopingCart
	Order
	Logger
}

func NewService(repos *repository.Repository, infoLogFile *os.File, errLogFile *os.File) *Service {
	return &Service{
		Users: NewUserService(repos.UsersRepo),
		Couriers: NewCourierService(repos.CouriersRepo),
		CouriersOrders: NewCouriersOrdersService(repos.CouriersOrdersRepo),
		Products: NewProductsService(repos.ProductsRepo),
		ShopingCart: NewShopingCartService(repos.ShopingCartRepo),
		Order: NewOrderService(repos.OrderRepo),
		Logger: NewServiceLogger(infoLogFile, errLogFile),
	}
}