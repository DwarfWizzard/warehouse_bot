package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) commandStart(message *tgbotapi.Message) error {
	if err := b.services.Create(message.Chat.ID); err != nil {
		return err
	}

	user, err := b.services.GetUser(message.Chat.ID)
	if err != nil {
		return err
	}

	if user.DialogueStatus == "pre_registration" {
		err := b.messagePreRegistration(message)
		if err != nil {
			return err
		}
	}
	return nil
}