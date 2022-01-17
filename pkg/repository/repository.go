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

type Repository struct {
	UsersRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UsersRepo: NewUsersSQLite3(db),
	}
}
