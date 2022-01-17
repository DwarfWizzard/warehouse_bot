package service

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type Users interface {
	Create(telegramId string) error
	GetUser(telegramId string) (models.User, error)
	UpdateUserName(telegramId string, name string) error
	UpdateUserNumber(telegramId string, number string) error
}

type Service struct {
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users: NewUserService(repos.UsersRepo),
	}
}