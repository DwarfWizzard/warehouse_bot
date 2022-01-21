package models

type User struct {
	Id             int    `db:"id"`
	TelegramId     string `db:"telegram_id"`
	Name           string `db:"name"`
	Number         string `db:"number"`
	DialogueStatus string `db:"dialogue_status"`
}
