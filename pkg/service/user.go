package service

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type UserService struct {
	repo repository.UsersRepo
}

func NewUserService(repo repository.UsersRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(telegramId int64) error {
	return s.repo.Create(telegramId)
}

func (s *UserService) GetUser(telegramId int64) (models.User, error) {
	user, err := s.repo.GetUser(telegramId)
	return user, err
}

func (s *UserService) UpdateUser(telegramId int64, field string, value string) error {
	return s.repo.UpdateUser(telegramId, field, value)
}

func (s *UserService) DeleteUser(telegramId int64) error {
	return s.repo.DeleteUser(telegramId)
}