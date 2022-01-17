package repository

import (
	"database/sql"
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
)
type UsersSQLite3 struct {
	db *sql.DB
}

func NewUsersSQLite3(db *sql.DB) *UsersSQLite3 {
	return &UsersSQLite3{
		db: db,
	}
}

func (r *UsersSQLite3) Create(telegramId int64) error {
	query := fmt.Sprintf("INSERT INTO %s (telegram_id) VALUES ($1)", usersTable)
	_, err := r.db.Exec(query, telegramId)

	return err
}

func (r *UsersSQLite3) GetUser(telegramId int64) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE telegram_id=$1", usersTable)
	row := r.db.QueryRow(query, telegramId)
	
	err := row.Scan(&user.Id, &user.TelegramId, &user.Name, &user.Number, &user.DialogueStatus)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UsersSQLite3) UpdateUserName(telegramId int64, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE telegram_id=$2", usersTable)
	_, err := r.db.Exec(query, name, telegramId)

	return err
}

func (r *UsersSQLite3) UpdateUserNumber(telegramId int64, number string) error {
	query := fmt.Sprintf("UPDATE %s SET number=$1 WHERE telegram_id=$2", usersTable)
	_, err := r.db.Exec(query, number, telegramId)

	return err
}

func (r *UsersSQLite3) UpdateUserStatus(telegramId int64, status string) error {
	query := fmt.Sprintf("UPDATE %s SET dialogue_status=$1 WHERE telegram_id=$2", usersTable)
	_, err := r.db.Exec(query, status, telegramId)

	return err
}