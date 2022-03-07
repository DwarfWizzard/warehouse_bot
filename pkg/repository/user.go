package repository

import (
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)
type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{
		db: db,
	}
}

func (r *UsersPostgres) Create(telegramId int64) error {
	query := fmt.Sprintf("INSERT INTO %s (telegram_id) VALUES ($1)", usersTable)
	_, err := r.db.Exec(query, telegramId)

	if err != nil {
		return fmt.Errorf("repository/CreateUser: [telegramId %d] : error %s", telegramId, err.Error())
	}

	return nil
}

func (r *UsersPostgres) GetUser(telegramId int64) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE telegram_id=$1", usersTable)
	err := r.db.Get(&user, query, telegramId)
	if err != nil {
		return models.User{}, fmt.Errorf("repository/GetUser: [telegramId %d] : error %s", telegramId, err.Error())
	}

	return user, nil
}

func (r *UsersPostgres) UpdateUser(telegramId int64, field string, value string) error {
	query := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE telegram_id=$2", usersTable, field)
	_, err := r.db.Exec(query, value, telegramId)

	if err != nil {
		return fmt.Errorf("repository/UpdateUser: [telegramId %d] [filed %s] [value %s] : %s", telegramId, field, value, err.Error())
	}

	return nil
}

func (r *UsersPostgres) DeleteUser(telegramId int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE telegram_id=$1", usersTable)
	_, err := r.db.Exec(query, telegramId)

	if err != nil {
		return fmt.Errorf("repository/CreateUser: [telegramId %d] : error %s", telegramId, err.Error())
	}

	return nil
}