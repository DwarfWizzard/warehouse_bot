package service

import (
	"time"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
)

type OrderService struct {
	repo repository.OrderRepo
}

func NewOrderService(repo repository.OrderRepo) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) GetOrderById(orderId int) (models.Order, error) {
	return s.repo.GetOrderById(orderId)
}

func (s *OrderService) GetOrderByUser(telegramId int64) (models.Order, error) {
	var order models.Order
	order, err := s.repo.GetOrderByUser(telegramId)
	if err != nil && (err.Error() == "sql: no rows in result set" || order.Id == 0) {
		date := time.Now().Format("02.01.2006 15:04:05")
		orderCreate, err := s.repo.Create(telegramId, date)

		if err != nil {
			return orderCreate, err
		}

		order = orderCreate
	} else if err != nil {
		return order, err
	}

	return order, nil
}

func (s *OrderService) GetAllOrdersUser(userId int) ([]models.Order, error) {
	return s.repo.GetAllOrdersUser(userId)
}

func (s *OrderService) UpdateOrder(telegramId int64, field string, value string) error {
	return s.repo.UpdateOrder(telegramId, field, value)
}

func (s *OrderService) GetOrderUser(orderId int) (models.User, error) {
	return s.repo.GetOrderUser(orderId)
}