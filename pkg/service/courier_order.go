package service

import (
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
	return s.repo.Create(orderId, courierId)
}

func (s *CourierOrderService) GetOrderCourier(orderId int) (models.Courier, error) {
	return s.repo.GetCourier(orderId)
}

func (s *CourierOrderService) GetActiveOrders(courierId int) ([]models.Order, error) {
	return s.repo.GetActiveOrders(courierId)
}

func (s *CourierOrderService) GetOrders(courierId int) ([]models.Order, error) {
	return s.repo.GetOrders(courierId)
}

func (s *CourierOrderService) UpdateCourierOrder(orderId int, field string, value string) error {
	return s.repo.Update(orderId, field, value)
}

func (s *CourierOrderService) GetOrderStatus(orderId int) (string, error) {
	return s.repo.GetOrderStatus(orderId)
}