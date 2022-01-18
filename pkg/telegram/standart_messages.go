package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) standartMessageRegistrationName(message *tgbotapi.Message) error{
	err := b.services.UpdateUserName(message.Chat.ID, message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateUserStatus(message.Chat.ID, "registration_add-number")
	if err != nil {
		return err
	}

	err = b.sendMessage(message.Chat.ID, "Отлично, теперь укажите ваш номер телофна.")
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) standartMessageRegistrationNumber(message *tgbotapi.Message) error{
	err := b.services.UpdateUserNumber(message.Chat.ID, message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateUserStatus(message.Chat.ID, "normal")
	if err != nil {
		return err
	}

	err = b.messageRegistrationLast(message)
	if err != nil {
		return err
	}
	return nil
}