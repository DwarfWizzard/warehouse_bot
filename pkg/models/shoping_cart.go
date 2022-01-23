package models

type ShopingCart struct {
	OrderId   int    `db:"order_id"`
	ProductId int    `db:"product_id"`
	Price     int `db:"price"`
	Quantity  int    `db:"quantity"`
}
