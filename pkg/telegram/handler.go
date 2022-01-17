package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.startCommand(message)
	}
	return nil

}

func (b *Bot) startCommand(message *tgbotapi.Message) error {
	return b.services.Create(message.Chat.ID)
}