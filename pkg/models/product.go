package models

type Product struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Price       string `db:"price"`
	Description string `db:"description"`
	Quantity    int
}
