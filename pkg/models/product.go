package models

type Product struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	PriceKilo   int    `db:"price_kg"`
	PriceBag    int    `db:"price_bag"`
	Description string `db:"description"`
	ImageName   string `db:"image_name"`
}
