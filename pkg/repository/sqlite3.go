package repository

import (
	"database/sql"
)

const (
	usersTable = "users"
	productsTable = "products"
	shopingCartTable = "shoping_cart"
	productListsTable = "product_lists"
)

func NewPostgresDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}