package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) sendMessage(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) sendMessages(args ...tgbotapi.MessageConfig) {
	for _, arg := range args {
		_, err := b.bot.Send(arg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func (b *Bot) sendMessageWithKeyboard(chatId int64, text string, Keyboard interface{}) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = Keyboard

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) updateMessage(chatId int64, messageId int, keyboard *tgbotapi.InlineKeyboardMarkup, newText string) error {
 	editMsg := tgbotapi.NewEditMessageText(chatId, messageId, newText)
	editMsg.ReplyMarkup = keyboard
	_, err := b.bot.Send(editMsg)

	return err
}

func (b *Bot) messageRegistrationLast(chatId int64) error {
	user, err := b.services.GetUser(chatId)
	if err != nil {
		return err
	}

	err = b.sendMessageWithKeyboard(chatId, fmt.Sprintf("Проверьте верность указанных данных. \nИмя:%s\nНомер:%s", user.Name, user.Number), registrationLastBoard)
	return err
}
