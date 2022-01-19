package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) deleteReplyMarkup(message *tgbotapi.Message) error {
	editMsg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0)})
	_, err := b.bot.Send(editMsg)

	return err
}

func (b *Bot) deleteMessage(chatID int64, messageId int) error {
	deleteMessage := tgbotapi.NewDeleteMessage(chatID, messageId)
	_, err := b.bot.Request(deleteMessage)

	return err
}
