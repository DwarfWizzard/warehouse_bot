package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) commandStart(message *tgbotapi.Message) error {
	_, err := b.services.GetUser(message.Chat.ID)

	if err != nil && strings.Contains(err.Error(), "sql: no rows in result set") {
		if err := b.services.CreateUser(message.Chat.ID); err != nil {
			return err
		}
		err := b.sendMessageWithKeyboard(message.Chat.ID, "Здравствуйте! С вами на связи бот торговой базы. \nВас я вижу первый раз. \nНажмите на зеленую кнопку, чтобы начать процесс регистрации.",
			preRegistrationBoard)

		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	
	return nil
}
