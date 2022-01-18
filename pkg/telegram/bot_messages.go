package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) sendMessage(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := b.bot.Send(msg)
	return err
}

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

var preRegistrationBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "pre_reg_yes"),
		tgbotapi.NewInlineKeyboardButtonData("🔴", "pre_reg_no"),
	),
)

func (b *Bot) messagePreRegistration(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Здравствуйте! С вами на связи бот торговой базы. \nВас я вижу первый раз. \nНажмите на зеленую кнопку, если вы хотите сохранить свой номер телефона и ваше имя для оформления заказов. Если вы отказываетесь, то при оформлении каждого заказа вам нужно будет указывать эти данные.")
	msg.ReplyMarkup = preRegistrationBoard

	_, err := b.bot.Send(msg)
	return err
}

var registrationLastBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "reg_last_yes"),
		tgbotapi.NewInlineKeyboardButtonData("🔴", "reg_last_no"),
	),
)

func (b *Bot) messageRegistrationLast(message *tgbotapi.Message) error {
	user, err := b.services.GetUser(message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,  fmt.Sprintf("Проверьте верность указанных данных. \nИмя:%s\nНомер:%s", user.Name, user.Number))
	msg.ReplyMarkup = registrationLastBoard

	_, err = b.bot.Send(msg)
	return err
}
