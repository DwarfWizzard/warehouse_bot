package repository

import (
	"errors"
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{
		db: db,
	}
}

func (r *OrderPostgres) Create(telegramId int64, date string) (models.Order, error) {
	var order models.Order
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE telegram_id=$1", usersTable)
	err := r.db.Get(&user, query, telegramId)
	if err != nil {
		return order, errors.New("repository/CreateOrder: select from users : "+err.Error())
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, user_name, user_number, user_city, order_date) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, user_name, user_number, order_date", ordersTable)
	err = r.db.Get(&order, query, user.Id, user.Name, user.Number, user.City, date)
	if err != nil {
		return order, fmt.Errorf("repository/CreateOrder: [telegramId %d] [date %s] : error %s", telegramId, date, err.Error())
	}
	
	return order, nil
}

func (r *OrderPostgres) GetOrderById(orderId int) (models.Order, error) {
	var order models.Order

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", ordersTable)
	err := r.db.Get(&order, query, orderId)
	if err != nil {
		return order, fmt.Errorf("repository/GetOrderById: [orderId %d] : error %s", orderId, err.Error())
	}

	return order, nil
}

func (r *OrderPostgres) GetOrderByUser(telegramId int64) (models.Order, error) {
	var order models.Order
	var userId int

	query := fmt.Sprintf("SELECT id FROM %s WHERE telegram_id=$1", usersTable)
	err := r.db.Get(&userId, query, telegramId)
	if err != nil {
		return order, fmt.Errorf("repository/GetOrder: [telegramId %d] : error %s", telegramId, err.Error())
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND order_status='in_progress' ", ordersTable)
	err = r.db.Get(&order, query, userId)
	if err != nil {
		return order, fmt.Errorf("repository/GetOrder: [user_id %d] : error %s",userId, err.Error())
	}

	return order, nil
}

func (r *OrderPostgres) GetOrderUser(orderId int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT tl.id, tl.telegram_id, tl.name, tl.number, tl.dialogue_status FROM %s tl INNER JOIN %s ul ON ul.user_id=tl.id WHERE ul.id=$1", usersTable, ordersTable)
	err := r.db.Get(&user, query, orderId)
	if err != nil {
		return user, fmt.Errorf("repository/GetOrderUser: [orderId %d] : error %s", orderId, err.Error())
	}

	return user, nil
}

func (r *OrderPostgres) UpdateOrder(telegramId int64, field string, value string) error {
	var userId int

	query := fmt.Sprintf("SELECT id FROM %s WHERE telegram_id=$1", usersTable)
	err := r.db.Get(&userId, query, telegramId)
	if err != nil {
		return fmt.Errorf("repository/UpdateOrder: [telegramId %d]  : error %s", telegramId,  err.Error())
	}

	query = fmt.Sprintf("UPDATE %s SET %s=$1 WHERE user_id=$2 AND order_status='in_progress'", ordersTable, field)
	_, err = r.db.Exec(query, value, userId)
	if err != nil {
		return fmt.Errorf("repository/UpdateOrder: [field %s] [userId %d] [value %s] : error %s", field, userId,value, err.Error())
	}

	return nil
}

func (r *OrderPostgres) GetAllOrdersUser(userId int) ([]models.Order, error) {
	var orders []models.Order

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", ordersTable)
	err := r.db.Select(&orders, query, userId)
	if err != nil {
		return orders, fmt.Errorf("repository/GetAllOrdersUser: [user_id %d] : error %s", userId, err.Error())
	}

	return orders, nil
}

