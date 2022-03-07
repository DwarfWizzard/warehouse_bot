package models

type Order struct {
	Id             int    `db:"id"`
	UserId         int    `db:"user_id"`
	UserName       string `db:"user_name"`
	UserNumber     string `db:"user_number"`
	UserCity       string `db:"user_city"`
	DeliveryAdress string `db:"delivery_adress"`
	OrderDate      string `db:"order_date"`
	OrderStatus    string `db:"order_status"`
}
