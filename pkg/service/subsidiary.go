package service

import (
	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type SubsidaryService struct {
	repo repository.SubsidiaryRepo
}

func NewSubsidaryService(repo repository.SubsidiaryRepo) *SubsidaryService {
	return &SubsidaryService{
		repo: repo,
	}
}

func (s *SubsidaryService) GetSubsidiary(cityName string) (models.Subsidiary, error) {
	return s.repo.GetSubsidiary(cityName)
} 

func (s *SubsidaryService) GetSubsidiarys() ([]models.Subsidiary, error) {
	return s.repo.GetSubsidiarys()
}