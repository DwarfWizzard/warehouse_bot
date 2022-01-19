package telegram

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) callbackPreReg(callbackQuery *tgbotapi.CallbackQuery) error {

	err := b.deleteReplyMarkup(callbackQuery.Message)
	if err != nil {
		return err
	}

	if callbackQuery.Data == callbackPreRegYes {
		err := b.services.UpdateUserStatus(callbackQuery.From.ID, "registration_add-name")
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

func (b *Bot) callbackChangePage(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	cbSplit := strings.Split(callbackQuery.Data, " ")
	page, err := strconv.Atoi(cbSplit[1])
	log.Println(page)
	
	if err != nil {
		return err
	}

	offset := (page-1) * 5
	log.Print(offset)

	prodOnPage, err := b.services.CountProductsOnPage(offset)
	log.Print(prodOnPage)

	if err != nil {
		return err
	}

	for i := 0; i <= prodOnPage; i++ {
		b.deleteMessage(chatId, callbackQuery.Message.MessageID-i)
	}


	productsList, err := b.generateProductsCardMessages(chatId, page)
	if err != nil {
		return err
	}

	b.sendMessages(productsList...)

	return nil
}
