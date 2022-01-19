package repository

import (
	"database/sql"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
)

type UsersRepo interface {
	Create(telegramId int64) error
	GetUser(telegramId int64) (models.User, error)
	UpdateUserName(telegramId int64, name string) error
	UpdateUserNumber(telegramId int64, number string) error
	UpdateUserStatus(telegramId int64, status string) error
}

type ProductsRepo interface {
	GetProducts(offset int) ([]models.Product, error)
	CountAllProducts() (int, error)
	CountProductsOnPage(offset int) (int, error) 
}

type ShopingCartRepo interface {
	Create(userId int) error
	Get(userId int) (models.Cart, error)
	Update(userId int, newCart models.Cart) error
	Delete(userId int) error
 }

type Repository struct {
	UsersRepo
	ProductsRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UsersRepo: NewUsersSQLite3(db),
		ProductsRepo: NewProductsSQLite3(db),
	}
}
