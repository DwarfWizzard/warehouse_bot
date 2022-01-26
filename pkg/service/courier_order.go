package service

import (
	"strings"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type CourierOrderService struct{
	repo repository.CouriersOrdersRepo
}

func NewCouriersOrdersService(repo repository.CouriersOrdersRepo) *CourierOrderService {
	return &CourierOrderService{
		repo: repo,
	}
}

func (s *CourierOrderService) CreateCourierOrder(orderId int, courierId int) error {
	_, err := s.repo.GetCourier(orderId)
	if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set") {
		return err
	} else {
		err := s.repo.Create(orderId, courierId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CourierOrderService) GetOrderCourier(orderId int) (models.Courier, error) {
	return s.repo.GetCourier(orderId)
}

func (s *CourierOrderService) GetActiveOrders(courierId int) ([]models.Order, error) {
	return s.repo.GetActiveOrders(courierId)
}

func (s *CourierOrderService) GetCourierOrders(courierId int) ([]models.Order, error) {
	return s.repo.GetCourierOrders(courierId)
}

func (s *CourierOrderService) UpdateCourierOrder(orderId int, field string, value string) error {
	return s.repo.Update(orderId, field, value)
}