package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
