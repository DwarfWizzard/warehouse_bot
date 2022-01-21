package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) commandStart(message *tgbotapi.Message) error {
	if err := b.services.CreateUser(message.Chat.ID); err != nil {
		return err
	}

	user, err := b.services.GetUser(message.Chat.ID)
	if err != nil {
		return err
	}

	if user.DialogueStatus == "pre_registration" {
		err := b.sendMessageWithKeyboard(message.Chat.ID, "Здравствуйте! С вами на связи бот торговой базы. \nВас я вижу первый раз. \nНажмите на зеленую кнопку, если вы хотите сохранить свой номер телефона и ваше имя для оформления заказов. Если вы отказываетесь, то при оформлении каждого заказа вам нужно будет указывать эти данные.", 
		preRegistrationBoard)
		
		if err != nil {
			return err
		}
	}
	return nil
}
