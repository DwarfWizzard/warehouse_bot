package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
)

const (
	callbackPreRegYes = "pre_reg_yes"
	callbackPreRegNo =  "pre_reg_no"
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.commandStart(message)
	}
	return nil

}

func (b *Bot) handleCallbacks(callbackQuery *tgbotapi.CallbackQuery) error {

	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	b.bot.Request(callback)

	switch callbackQuery.Data {
	case callbackPreRegYes, callbackPreRegNo:
		
	}
	return nil
}

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

func (b *Bot) callbackPreReg(callbackQuery *tgbotapi.CallbackQuery) error {
	return nil
}