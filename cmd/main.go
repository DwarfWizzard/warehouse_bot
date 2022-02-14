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
	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	year, month, day := time.Now().Date()

	errLogFile, err := os.OpenFile("logs/"+fmt.Sprintf("err_%v-%v-%v.log", year, month, day), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("run error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, errLogFile)

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("failied to init telegram bot: %s", err.Error())
	}

	botApi.Debug=false

	bot := telegram.NewBot(botApi, services)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}