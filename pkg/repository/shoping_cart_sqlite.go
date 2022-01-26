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
		return fmt.Errorf("repository/CreateShopingCart: [productId %d] : error %s", productId, err.Error())
	}

	query = fmt.Sprintf("INSERT INTO %s (order_id, product_id, price) VALUES ($1, $2, $3)", shopingCartTable)
	_, err = r.db.Exec(query, orderId, productId, product.Price)

	if err != nil {
		return fmt.Errorf("repository/CreateShopingCart: [orderId %d] [productId %d] : error %s", orderId, productId, err.Error())
	}
	return nil
}

func (r *ShopingCartSQLite3) GetCart(orderId int, productId int) (models.ShopingCart, error) {
	var cart models.ShopingCart
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	err := r.db.Get(&cart, query, orderId, productId)

	if err != nil {
		return cart, fmt.Errorf("repository/GetCart: [orderId %d] [productId %d] : error %s", orderId, productId, err.Error())
	}

	return cart, nil
}

func (r *ShopingCartSQLite3) GetProductsFromCart(orderId int) ([]models.Product, error) {
	var products []models.Product
	
	query := fmt.Sprintf("SELECT tl.id, tl.title, ul.price, tl.description, tl.image_name FROM %s tl INNER JOIN %s ul WHERE ul.order_id=$1 AND ul.product_id=tl.id;", productsTable, shopingCartTable)
	err := r.db.Select(&products, query, orderId)
	if err != nil {
		return products, fmt.Errorf("repository/GetProductsFromCart: [orderId %d] : error %s", orderId, err.Error())
	}
	
	return products, nil
}

func (r *ShopingCartSQLite3) GetQuantity(orderId int, productId int) (int, error) {
	var quantity int

	query := fmt.Sprintf("SELECT quantity FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	err := r.db.Get(&quantity, query, orderId, productId)

	if err != nil {
		return 0, fmt.Errorf("repository/GetQuantity: [orderId %d] [productId %d] : error %s", orderId, productId, err.Error())
	}

	return quantity, nil
}

func (r *ShopingCartSQLite3) UpdateQuantity(orderId int, productId int, quantity int) error {
	query := fmt.Sprintf("UPDATE %s SET quantity=$1, price=(SELECT price FROM %s WHERE id=$2)*$1 WHERE order_id=$3 AND product_id=$4", shopingCartTable, productsTable)
	_, err := r.db.Exec(query, quantity, productId, orderId, productId)

	if err != nil {
		return fmt.Errorf("repository/UpdateQuantity: [orderId %d] [productId %d] [quantity %d]: error %s", orderId, productId, quantity, err.Error())
	}

	return nil
}

func (r *ShopingCartSQLite3) DeleteCart(orderId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1", shopingCartTable)
	_, err := r.db.Exec(query, orderId)
	
	if err != nil {
		return fmt.Errorf("repository/DeleteCart: [orderId %d] : error %s", orderId, err.Error())
	}

	return nil
}

func (r *ShopingCartSQLite3) DeleteProductFromCart(orderId int, productId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	_, err := r.db.Exec(query, orderId, productId)

	if err != nil {
		return fmt.Errorf("repository/DeleteProductFromCart: [orderId %d] [productId %d] : error %s", orderId, productId, err.Error())
	}

	return err
}