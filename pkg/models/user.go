package models

type User struct {
	Id             int    `db:"id"`
	TelegramId     int64 `db:"telegram_id"`
	Name           string `db:"name"`
	Number         string `db:"number"`
	DialogueStatus string `db:"dialogue_status"`
}

type Courier struct {
	Id             int    `db:"id"`
	TelegramId     int64 `db:"telegram_id"`
	Name           string `db:"name"`
	Number         string `db:"number"`
	DialogueStatus string `db:"dialogue_status"`
}
