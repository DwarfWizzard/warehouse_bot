package repository

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"log"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
)

type ProductsSQLite3 struct {
	db *sqlx.DB
}

func NewProductsSQLite3(db *sqlx.DB) *ProductsSQLite3 {
	return &ProductsSQLite3{
		db: db,
	}
}

func (r *ProductsSQLite3) GetProducts(offset int) ([]models.Product, error) {
	var products []models.Product

	query := fmt.Sprintf("SELECT * FROM %s LIMIT 5 OFFSET $2", productsTable)
	rows, err := r.db.Query(query, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.Id, &product.Title, &product.Price, &product.Description)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductsSQLite3) CountProducts() (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", productsTable)
	row := r.db.QueryRow(query)

	err := row.Scan(&count)
	return count, err
}

func (r *ProductsSQLite3) CountProductsOnPage(offset int) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM (SELECT * FROM %s LIMIT 5 OFFSET $2)", productsTable)
	row := r.db.QueryRow(query, offset)

	err := row.Scan(&count)
	return count, err
}
