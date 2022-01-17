package service

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type Users interface {
	Create(telegramId int64) error
	GetUser(telegramId int64) (models.User, error)
	UpdateUserName(telegramId int64, name string) error
	UpdateUserNumber(telegramId int64, number string) error
	UpdateUserStatus(telegramId int64, status string) error
}

type Service struct {
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users: NewUserService(repos.UsersRepo),
	}
}