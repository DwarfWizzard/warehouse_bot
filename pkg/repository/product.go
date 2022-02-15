package repository

import (
	"github.com/jmoiron/sqlx"
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
)

type ProductsPostgres struct {
	db *sqlx.DB
}

func NewProductsPostgres(db *sqlx.DB) *ProductsPostgres {
	return &ProductsPostgres{
		db: db,
	}
}

func (r *ProductsPostgres) GetProduct(productId int) (models.Product, error) {
	var product models.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productsTable)
	err := r.db.Get(&product, query, productId)

	if err != nil {
		return product, fmt.Errorf("repositoty/GetProduct: [productId %d] : error %s", productId, err.Error())
	}

	return product, nil
}

func (r *ProductsPostgres) GetProducts(offset int) ([]models.Product, error) {
	var products []models.Product

	query := fmt.Sprintf("SELECT * FROM %s LIMIT 5 OFFSET $1", productsTable)
	err := r.db.Select(&products, query, offset)
	if err != nil {
		return nil, fmt.Errorf("repositoty/GetProducts: [offset%d] : error %s", offset, err.Error())
	}

	return products, nil
}

func (r *ProductsPostgres) CountProducts() (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", productsTable)
	row := r.db.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return 0, fmt.Errorf("repositoty/CountProducts: error %s", err.Error())
	}

	return count, nil
}

func (r *ProductsPostgres) CountProductsOnPage(offset int) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM (SELECT * FROM %s LIMIT 5 OFFSET $2)", productsTable)
	row := r.db.QueryRow(query, offset)

	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("repositoty/CountProductsOnPage: [offset%d] : error %s", offset, err.Error())
	}

	return count, err
}
