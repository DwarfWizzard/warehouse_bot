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

func (s *UserService) Create(telegramId int64) error {
	return s.repo.Create(telegramId)
}

func (s *UserService) GetUser(telegramId int64) (models.User, error) {
	user, err := s.repo.GetUser(telegramId)
	
	return user, err
}

func (s *UserService) UpdateUserName(telegramId int64, name string) error {
	return s.repo.UpdateUserName(telegramId, name)
}
func (s *UserService) UpdateUserNumber(telegramId int64, number string) error {
	return s.repo.UpdateUserNumber(telegramId, number)
}

func (s *UserService) UpdateUserStatus(telegramId int64, status string) error {
	return s.repo.UpdateUserStatus(telegramId, status)
}