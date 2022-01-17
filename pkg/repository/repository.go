package repository

import (
	"database/sql"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
)

type UsersRepo interface {
	Create(telegramId string) error
	GetUser(telegramId string) (models.User, error)
	UpdateUserName(telegramId string, name string) error
	UpdateUserNumber(telegramId string, number string) error
}

type Repository struct {
	UsersRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UsersRepo: NewUsersSQLite3(db),
	}
}
