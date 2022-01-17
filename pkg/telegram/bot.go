package telegram

import (
	"log"

	"github.com/DwarfWizzard/warehouse_bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	services *service.Service
}

func NewBot(bot *tgbotapi.BotAPI, services *service.Service) *Bot {
	return &Bot{
		bot:      bot,
		services: services,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			if err := b.handleCommands(update.Message); err != nil {
				log.Print(err.Error())
			}
		} else if update.CallbackQuery != nil {
			log.Println("q11231")
			if err := b.handleCallbacks(update.CallbackQuery); err != nil {
				log.Print(err.Error())
			}
		}

	}
	return nil
}
