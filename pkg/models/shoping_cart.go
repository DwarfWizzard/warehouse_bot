package models

type ShopingCart struct {
	UserId int `db:"user_id"`
	ProductId int `db:"product_id"`
	Quantity int `db:"quantity"`
}