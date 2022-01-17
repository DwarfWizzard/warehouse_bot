package main

import (
	"log"
	"os"

	"github.com/DwarfWizzard/warehouse_bot/pkg/repository"
	"github.com/DwarfWizzard/warehouse_bot/pkg/service"
	"github.com/DwarfWizzard/warehouse_bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewSQLite3DB(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	key := os.Getenv("KEY")
	botApi, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Fatalf("failied to init telegram bot: %s", err.Error())
	}

	botApi.Debug=false

	bot := telegram.NewBot(botApi, services)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}