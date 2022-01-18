package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) callbackPreReg(callbackQuery *tgbotapi.CallbackQuery) error {

	err := b.deleteReplyMarkup(callbackQuery.Message)
	if err != nil {
		return err
	}

	if callbackQuery.Data == callbackPreRegYes {
		err := b.services.UpdateUserStatus(callbackQuery.Message.Chat.ID, "registration_add-name")
		if err != nil {
			return err
		}

		err = b.sendMessage(callbackQuery.From.ID, "Хорошо. Пожалуйста, напишите ваше имя.")
		if err != nil {
			return err
		}
	}
	if callbackQuery.Data == callbackPreRegNo {
		err := b.services.UpdateUserStatus(callbackQuery.From.ID, "normal")
		if err != nil {
			return err
		}

		err = b.sendMessage(callbackQuery.From.ID, "Хорошо. При оформлении заказов, вам необходимо будет указывать эти данные.")
		if err != nil {
			return err
		}
	}

	

	return nil
}

func (b *Bot) callbackRegLast(callbackQuery *tgbotapi.CallbackQuery) error {

	if callbackQuery.Data == callbackRegLastYes {
		err := b.sendMessage(callbackQuery.From.ID, "Замечательно! Приятно с вами познакомиться :)")
		if err != nil {
			return err
		}
	}
	if callbackQuery.Data == callbackRegLastNo {
		err := b.services.UpdateUserStatus(callbackQuery.From.ID, "registration_add-name")
		if err != nil {
			return err
		}

		err = b.sendMessage(callbackQuery.From.ID, "Хорошо. Пожалуйста, напишите ваше имя.")
		if err != nil {
			return err
		}
	}

	err := b.deleteReplyMarkup(callbackQuery.Message)
	if err != nil {
		return err
	}
	return nil
}
