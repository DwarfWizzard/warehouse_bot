package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) deleteReplyMarkup(message *tgbotapi.Message) error {
	editMsg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0)})
	_, err := b.bot.Send(editMsg)

	if err != nil {
		return fmt.Errorf("telegram/deleteReplyMarkup: error %s", err.Error())
	}

	return nil
}

func (b *Bot) deleteMessage(chatID int64, messageId int) error {
	deleteMessage := tgbotapi.NewDeleteMessage(chatID, messageId)
	_, err := b.bot.Request(deleteMessage)

	if err != nil {
		return fmt.Errorf("telegram/deleteMessage: error %s", err.Error())
	}

	return nil
}
