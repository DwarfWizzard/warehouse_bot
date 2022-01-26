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
	var product models.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productsTable)
	err := r.db.Get(&product, query, productId)
	if err != nil {
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (order_id, product_id, price) VALUES ($1, $2, $3)", shopingCartTable)
	_, err = r.db.Exec(query, orderId, productId, product.Price)

	return err
}

func (r *ShopingCartSQLite3) GetCart(orderId int, productId int) (models.ShopingCart, error) {
	var cart models.ShopingCart
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	err := r.db.Get(&cart, query, orderId, productId)

	return cart, err
}

func (r *ShopingCartSQLite3) GetProductsFromCart(orderId int) ([]models.Product, error) {
	var products []models.Product
	
	query := fmt.Sprintf("SELECT tl.id, tl.title, ul.price, tl.description, tl.image_name FROM %s tl INNER JOIN %s ul WHERE ul.order_id=$1 AND ul.product_id=tl.id;", productsTable, shopingCartTable)
	err := r.db.Select(&products, query, orderId)
	if err != nil {
		return products, err
	}
	
	return products, nil
}

func (r *ShopingCartSQLite3) GetQuantity(orderId int, productId int) (int, error) {
	var quantity int

	query := fmt.Sprintf("SELECT quantity FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	err := r.db.Get(&quantity, query, orderId, productId)

	return quantity, err
}

func (r *ShopingCartSQLite3) UpdateQuantity(orderId int, productId int, quantity int) error {
	query := fmt.Sprintf("UPDATE %s SET quantity=$1, price=(SELECT price FROM %s WHERE id=$2)*$1 WHERE order_id=$3 AND product_id=$4", shopingCartTable, productsTable)
	_, err := r.db.Exec(query, quantity, productId, orderId, productId)

	return err
}

func (r *ShopingCartSQLite3) DeleteCart(orderId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1", shopingCartTable)
	_, err := r.db.Exec(query, orderId)
	
	return err
}

func (r *ShopingCartSQLite3) DeleteProductFromCart(orderId int, productId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	_, err := r.db.Exec(query, orderId, productId)

	return err
}