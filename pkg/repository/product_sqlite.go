package repository

import (
	"github.com/jmoiron/sqlx"
	"fmt"

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

func (r *ProductsSQLite3) GetProduct(productId int) (models.Product, error) {
	var product models.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productsTable)
	err := r.db.Get(&product, query, productId)

	if err != nil {
		return product, fmt.Errorf("repositoty/GetProduct: [productId %d] : error %s", productId, err.Error())
	}

	return product, nil
}

func (r *ProductsSQLite3) GetProducts(offset int) ([]models.Product, error) {
	var products []models.Product

	query := fmt.Sprintf("SELECT * FROM %s LIMIT 5 OFFSET $2", productsTable)
	err := r.db.Select(&products, query, offset)
	if err != nil {
		return nil, fmt.Errorf("repositoty/GetProducts: [offset%d] : error %s", offset, err.Error())
	}

	return products, nil
}

func (r *ProductsSQLite3) CountProducts() (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", productsTable)
	row := r.db.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return 0, fmt.Errorf("repositoty/CountProducts: error %s", err.Error())
	}

	return count, nil
}

func (r *ProductsSQLite3) CountProductsOnPage(offset int) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM (SELECT * FROM %s LIMIT 5 OFFSET $2)", productsTable)
	row := r.db.QueryRow(query, offset)

	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("repositoty/CountProductsOnPage: [offset%d] : error %s", offset, err.Error())
	}

	return count, err
}
