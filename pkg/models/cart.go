package models

type Cart struct {
	Id           int    `db:"int"`
	UserId       int    `db:"user_id"`
	Adress       string `db:"adress"`
	DeliveryDate string `db:"delivery_date"`
}
