package repository

import (

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	productsTable = "products"
	shopingCartTable = "shoping_cart"
	ordersTable = "orders"
)

func NewSQLite3DB(path string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}