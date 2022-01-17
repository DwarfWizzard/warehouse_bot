package telegram

import (
	"log"

	"github.com/DwarfWizzard/warehouse_bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	services *service.Service
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot:   bot,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			if err := b.handleCommands(update.Message); err != nil {
				log.Print(err.Error())
			}
		}
		

	}
	return nil
}