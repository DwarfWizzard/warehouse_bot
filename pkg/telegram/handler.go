package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
)

const (
	callbackPreRegYes  = "pre_reg_yes"
	callbackPreRegNo   = "pre_reg_no"
	callbackRegLastYes = "reg_last_yes"
	callbackRegLastNo  = "reg_last_no"
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.commandStart(message)
	}
	return nil

}

func (b *Bot) handleCallbacks(callbackQuery *tgbotapi.CallbackQuery) error {
	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	b.bot.Request(callback)

	switch callbackQuery.Data {
	case callbackPreRegYes, callbackPreRegNo:
		err := b.callbackPreReg(callbackQuery)
		if err != nil {
			return err
		}
	case callbackRegLastYes, callbackRegLastNo:
		err := b.callbackRegLast(callbackQuery)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) handleStandartMessages(message *tgbotapi.Message) error {
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
