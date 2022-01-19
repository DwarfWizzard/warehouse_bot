package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) registration(message *tgbotapi.Message) error {
	user, err := b.services.GetUser(message.Chat.ID)
	if err != nil {
		return err
	}

	switch user.DialogueStatus {
	case "registration_add-name":
		err := b.standartMessageRegistrationName(message)
		if err != nil {
			return err
		}

	case "registration_add-number":
		err := b.standartMessageRegistrationNumber(message)
		if err != nil {
			return err
		}
	}

	return nil
}

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
	
	err = b.messageRegistrationLast(message.Chat.ID)
	if err != nil {
		return err
	}
	return nil
}


func (b *Bot) standartMessageCatalog(chatId int64) error {
	productsList, err := b.generateProductsCardMessages(chatId, 1)
	if err != nil {
		return err
	}

	b.sendMessages(productsList...)

	return nil
}