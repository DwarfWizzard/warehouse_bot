package repository

import (
	"database/sql"
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
)

type CartSQLite3 struct {
	db *sql.DB
}

func NewCartSQLite3(db *sql.DB) *CartSQLite3 {
	return &CartSQLite3{
		db: db,
	}
}

func (r *CartSQLite3) Create(userId int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id) VALUES ($1) RET", shopingCartTable)
	_, err := r.db.Exec(query, userId)

	return err
}

func (r *CartSQLite3) Get(userId int) (models.Cart, error) {
	var shopingCart models.Cart
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", shopingCartTable)

	row := r.db.QueryRow(query, userId)
	err := row.Scan(&shopingCart)

	return shopingCart, err
}

func (r *CartSQLite3) Update(userId int, newCart models.Cart) error {
	query := fmt.Sprintf("UPDATE %s SET adress=$1, delivery_date-$2 WHERE user_id=$3", shopingCartTable)
	_, err := r.db.Exec(query, newCart.Adress, newCart.DeliveryDate, userId)

	return err
}