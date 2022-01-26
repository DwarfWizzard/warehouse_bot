package repository

import (
	"errors"
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type CourierSQLite3 struct {
	db *sqlx.DB
}

func NewCourierSQLite3(db *sqlx.DB) *CourierSQLite3 {
	return &CourierSQLite3{
		db: db,
	}
}

func (r *CourierSQLite3) Create(telegramId int64) error {
	query := fmt.Sprintf("INSERT INTO %s (telegram_id) VALUES ($1)", couriersTable)
	_, err := r.db.Exec(query, telegramId)

	if err != nil {
		return errors.New("repository/CreateCourier: insert into couriers :"+err.Error())
	}

	return nil
}

func (r *CourierSQLite3) GetCourier(telegramId int64) (models.Courier, error) {
	var courier models.Courier

	query := fmt.Sprintf("SELECT * FROM %s WHERE telegram_id=$1", couriersTable)
	err := r.db.Get(&courier, query, telegramId)
	if err != nil {
		return courier, errors.New("repository/GetCourier: select from couriers : "+err.Error())
	}

	return courier, nil
}

func (r *CourierSQLite3) UpdateCourier(telegramId int64, field string, value string) error {
	query := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE telegram_id=$2", couriersTable, field)
	_, err := r.db.Exec(query, value, telegramId)
	if err != nil {
		return errors.New("repository/UpdateCourier: select from couriers : "+err.Error())
	}

	return nil
}