package repository

import (
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type ShopingCartSQLite3 struct {
	db *sqlx.DB
}

func NewShopingCartSQLite3(db *sqlx.DB) *ShopingCartSQLite3 {
	return &ShopingCartSQLite3{
		db: db,
	}
}

func (r *ShopingCartSQLite3) Create(orderId int, productId int) error {
	query := fmt.Sprintf("INSERT INTO %s (order_id, product_id) VALUES ($1, $2)", shopingCartTable)
	_, err := r.db.Exec(query, orderId, productId)

	return err
}

func (r *ShopingCartSQLite3) GetProducts(orderId int) ([]models.Product, error) {
	var products []models.Product
	
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.price, tl.description FROM %s tl INNER JOIN %s ul WHERE ul.order_id=$1 AND ul.product_id=tl.id;", productsTable, shopingCartTable)
	err := r.db.Select(&products, query, orderId)
	if err != nil {
		return products, err
	}

	return products, nil
}