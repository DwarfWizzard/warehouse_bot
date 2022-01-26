package repository

import (
	"errors"
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type OrderSQLite3 struct {
	db *sqlx.DB
}

func NewOrderSQLite3(db *sqlx.DB) *OrderSQLite3 {
	return &OrderSQLite3{
		db: db,
	}
}

func (r *OrderSQLite3) Create(telegramId int64, date string) (models.Order, error) {
	var order models.Order
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE telegram_id=$1", usersTable)
	err := r.db.Get(&user, query, telegramId)
	if err != nil {
		return order, errors.New("repository/CreateOrder: select from users : "+err.Error())
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, user_name, user_number, order_date) VALUES ($1, $2, $3, $4) RETURNING id, user_id, user_name, user_number, order_date", ordersTable)
	err = r.db.Get(&order, query, user.Id, user.Name, user.Number, date)
	if err != nil {
		return order, errors.New("repository/CreateOrder: insert into orders : "+err.Error())
	}
	
	return order, nil
}

func (r *OrderSQLite3) GetOrderById(orderId int) (models.Order, error) {
	var order models.Order

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", ordersTable)
	err := r.db.Get(&order, query, orderId)
	if err != nil {
		return order, errors.New("repository/GetOrderById: select from orders : "+err.Error())
	}

	return order, nil
}

func (r *OrderSQLite3) GetOrderByUser(telegramId int64) (models.Order, error) {
	var order models.Order
	var userId int

	query := fmt.Sprintf("SELECT id FROM %s WHERE telegram_id=$1", usersTable)
	err := r.db.Get(&userId, query, telegramId)
	if err != nil {
		return order, errors.New("repository/GetOrder: select from user : "+err.Error())
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND order_status=\"in_progress\"", ordersTable)
	err = r.db.Get(&order, query, userId)
	if err != nil {
		return order, errors.New("repository/GetOrder: select from orders : "+err.Error())
	}

	return order, nil
}

func (r *OrderSQLite3) GetOrderUser(orderId int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT tl.id, tl.telegram_id, tl.name, tl.number, tl.dialogue_status FROM %s tl INNER JOIN %s ul ON ul.user_id=tl.id WHERE ul.id=$1", usersTable, ordersTable)
	err := r.db.Get(&user, query, orderId)
	if err != nil {
		return user, errors.New("repository/GetOrderUser: "+err.Error())
	}

	return user, nil
}

func (r *OrderSQLite3) UpdateOrder(telegramId int64, field string, value string) error {
	var userId int

	query := fmt.Sprintf("SELECT id FROM %s WHERE telegram_id=$1", usersTable)
	err := r.db.Get(&userId, query, telegramId)
	if err != nil {
		return errors.New("repository/UpdateOrder: select from user :"+err.Error())
	}

	query = fmt.Sprintf("UPDATE %s SET %s=$1 WHERE user_id=$2 AND order_status=\"in_progress\"", ordersTable, field)
	_, err = r.db.Exec(query, value, userId)
	if err != nil {
		return errors.New("repository/UpdateOrder: update order :"+err.Error())
	}

	return nil
}

