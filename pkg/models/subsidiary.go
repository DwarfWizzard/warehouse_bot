package models

type Subsidiary struct {
	Id     int    `db:"id"`
	City   string `db:"city"`
	ChatId int64  `db:"chat_id"`
}
