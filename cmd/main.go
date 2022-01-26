package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	year, month, day := time.Now().Date()

	infoLogFile, err := os.OpenFile("logs/"+fmt.Sprintf("info_%v-%v-%v.log", year, month, day), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("run error: %s", err.Error())
	}
	errLogFile, err := os.OpenFile("logs/"+fmt.Sprintf("err_%v-%v-%v.log", year, month, day), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("run error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, infoLogFile, errLogFile)

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