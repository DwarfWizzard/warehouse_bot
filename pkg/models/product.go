package models

type Product struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Price       int `db:"price"`
	Description string `db:"description"`
	Quantity    int
}
