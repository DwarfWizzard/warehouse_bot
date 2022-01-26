package service

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type CourierService struct {
	repo repository.CouriersRepo
}

func NewCourierService(repo repository.CouriersRepo) *CourierService {
	return &CourierService{
		repo: repo,
	}
}

func (s *CourierService) CreateCourier(telegramId int64) error {
	return s.repo.Create(telegramId)
}

func (s *CourierService) GetCourier(telegramId int64) (models.Courier, error) {
	user, err := s.repo.GetCourier(telegramId)
	return user, err
}

func (s *CourierService) UpdateCourier(telegramId int64, field string, value string) error {
	return s.repo.UpdateCourier(telegramId, field, value)
}
