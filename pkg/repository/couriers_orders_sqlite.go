package repository

import (
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type CouriersOrdersSQLite3 struct {
	db *sqlx.DB
}

func NewCouriersOrdersSQLite3(db *sqlx.DB) *CouriersOrdersSQLite3 {
	return &CouriersOrdersSQLite3{
		db: db,
	}
}

func (r *CouriersOrdersSQLite3) Create(orderId int, courierId int) error {
	query := fmt.Sprintf("INSERT INTO %s (courier_id, order_id) VALUES ($1, $2)", couriersOrdersTable)
	_, err := r.db.Exec(query, courierId, orderId)
	if err != nil {
		return fmt.Errorf("repository/CreateCourirersOrders: [orderId %d] [courierId %d] : error %s", orderId, courierId, err.Error())
	}

	return nil
}

func (r *CouriersOrdersSQLite3) GetCourier(orderId int) (models.Courier, error) {
	var courier models.Courier
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1", couriersOrdersTable)
	err := r.db.Get(&courier, query, orderId)
	if err != nil {
		return courier, fmt.Errorf("repository/GetCourierByOrder: [orderId %d] : error %s", orderId, err.Error())
	}

	return courier, nil
}

func (r *CouriersOrdersSQLite3) GetActiveOrders(courierId int) ([]models.Order, error) {
	var orders []models.Order
	query := fmt.Sprintf("SELECT tl.id, tl.user_id, tl.user_name, tl.user_number, tl.delivery_adress, tl.order_date, tl.order_status FROM %s tl INNER JOIN %s ul ON ul.order_id=tl.id WHERE ul.status=\"active\" AND ul.courier_id=$1", ordersTable,couriersOrdersTable)
	err := r.db.Select(&orders, query, courierId)
	if err != nil {
		return orders, fmt.Errorf("repository/GetActiveOrders: [courierId %d] : error %s", courierId, err.Error())
	}

	return orders, nil
}

func (r *CouriersOrdersSQLite3) GetOrders(courierId int) ([]models.Order, error) {
	var orders []models.Order
	query := fmt.Sprintf("SELECT tl.id, tl.user_id, tl.user_name, tl.user_number, tl.delivery_adress, tl.order_date, tl.order_status FROM %s tl INNER JOIN %s ul ON ul.order_id=tl.id WHERE ul.courier_id=$1", ordersTable,couriersOrdersTable)
	err := r.db.Select(&orders, query, courierId)
	if err != nil {
		return orders, fmt.Errorf("repository/GetOrdersByCourier: [courierId %d] : error %s", courierId, err.Error())
	}

	return orders, nil
}

func (r *CouriersOrdersSQLite3) GetOrderStatus(orderId int) (string, error) {
	var status string
	query := fmt.Sprintf("SELECT status FROM %s WHERE order_id=$1", couriersOrdersTable)
	err := r.db.Get(&status, query, orderId)
	if err != nil {
		return status, fmt.Errorf("repository/GetOrdersStatus: [orderId %d] : error %s", orderId, err.Error())
	}

	return status, nil
}

func (r *CouriersOrdersSQLite3) Update(orderId int, field string, value string) error {
	query := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE order_id=$2 AND status=\"active\"", couriersOrdersTable, field)
	_, err := r.db.Exec(query, value, orderId)
	if err != nil {
		return fmt.Errorf("repository/UpdateCouriersOrders: [orderId %d] [filed %s] [value %s] : %s",orderId, field, value, err.Error())
	}

	return nil
}