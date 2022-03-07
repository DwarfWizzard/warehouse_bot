package telegram

import (
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
		if update.CallbackQuery != nil {
			if err := b.handleCallbacks(update.CallbackQuery); err != nil {
				b.services.PrintLog(err.Error(), 1)
			}
		} else if update.Message != nil && update.Message.IsCommand() && update.Message.Chat.ID > 1 {
			if err := b.handleCommands(update.Message); err != nil {
				b.services.PrintLog(err.Error(), 1)
			}
		} else if update.Message != nil && update.Message.Chat.ID > 1 {
			if err := b.handleStandartMessages(update.Message); err != nil {
				b.services.PrintLog(err.Error(), 1)
			}
		}
	}
	
	return nil
}
