package repository

import (
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type ShopingCartPostgres struct {
	db *sqlx.DB
}

func NewShopingCartPostgres(db *sqlx.DB) *ShopingCartPostgres {
	return &ShopingCartPostgres{
		db: db,
	}
}

func (r *ShopingCartPostgres) Create(orderId int, productId int, productPrice int, deliveryFormat string) error {
	query := fmt.Sprintf("INSERT INTO %s (order_id, product_id, price, unit_price, delivery_format) VALUES ($1, $2, $3, $4, $5)", shopingCartTable)
	_, err := r.db.Exec(query, orderId, productId, productPrice, productPrice, deliveryFormat)

	if err != nil {
		return fmt.Errorf("repository/CreateShopingCart: [orderId %d] [productId %d] : error %s", orderId, productId, err.Error())
	}
	return nil
}

func (r *ShopingCartPostgres) GetCart(orderId int, productId int) (models.ShopingCart, error) {
	var cart models.ShopingCart
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	err := r.db.Get(&cart, query, orderId, productId)

	if err != nil {
		return cart, fmt.Errorf("repository/GetCart: [orderId %d] [productId %d] : error %s", orderId, productId, err.Error())
	}

	return cart, nil
}

func (r *ShopingCartPostgres) GetCarts(orderId int) ([]models.ShopingCart, error) {
	var carts []models.ShopingCart
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_id=$1", shopingCartTable)
	err := r.db.Select(&carts, query, orderId)

	if err != nil {
		return carts, fmt.Errorf("repository/GetCarts: [orderId %d] : error %s", orderId, err.Error())
	}

	return carts, nil
}

func (r *ShopingCartPostgres) GetProductsFromCart(orderId int) ([]models.Product, error) {
	var products []models.Product
	
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.price_kg, tl.price_bag, tl.description, tl.image_name FROM %s tl INNER JOIN %s ul ON ul.product_id=tl.id WHERE ul.order_id=$1 AND ul.product_id=tl.id;", productsTable, shopingCartTable)
	err := r.db.Select(&products, query, orderId)
	if err != nil {
		return products, fmt.Errorf("repository/GetProductsFromCart: [orderId %d] : error %s", orderId, err.Error())
	}
	
	return products, nil
}

func (r *ShopingCartPostgres) GetQuantity(orderId int, productId int) (int, error) {
	var quantity int

	query := fmt.Sprintf("SELECT quantity FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	err := r.db.Get(&quantity, query, orderId, productId)

	if err != nil {
		return 0, fmt.Errorf("repository/GetQuantity: [orderId %d] [productId %d] : error %s", orderId, productId, err.Error())
	}

	return quantity, nil
}

func (r *ShopingCartPostgres) UpdateQuantity(orderId int, productId int, quantity int) error {
	query := fmt.Sprintf("UPDATE %s SET quantity=$1, price=unit_price*$1 WHERE order_id=$2 AND product_id=$3", shopingCartTable)
	_, err := r.db.Exec(query, quantity, orderId, productId)

	if err != nil {
		return fmt.Errorf("repository/UpdateQuantity: [orderId %d] [productId %d] [quantity %d]: error %s", orderId, productId, quantity, err.Error())
	}

	return nil
}

func (r *ShopingCartPostgres) DeleteCart(orderId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1", shopingCartTable)
	_, err := r.db.Exec(query, orderId)
	
	if err != nil {
		return fmt.Errorf("repository/DeleteCart: [orderId %d] : error %s", orderId, err.Error())
	}

	return nil
}

func (r *ShopingCartPostgres) DeleteProductFromCart(orderId int, productId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_id=$1 AND product_id=$2", shopingCartTable)
	_, err := r.db.Exec(query, orderId, productId)

	if err != nil {
		return fmt.Errorf("repository/DeleteProductFromCart: [orderId %d] [productId %d] : error %s", orderId, productId, err.Error())
	}

	return err
}