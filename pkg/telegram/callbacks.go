package telegram

import (
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
		err := b.sendMessageWithKeyboard(callbackQuery.From.ID, "Замечательно! Приятно с вами познакомиться :)", menuKeyboard)
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
	if err != nil {
		return err
	}

	offset := (page-1) * 5

	products, err := b.services.GetProducts(offset)
	if err != nil {
		return err
	}

	prodOnPage, err := b.services.CountProductsOnPage((page-2)*5)

	if err != nil {
		return err
	}

	for i := 0; i <= prodOnPage; i++ {
		b.deleteMessage(chatId, callbackQuery.Message.MessageID-i)
	}


	productsList := b.generateProductListCardsMessages(products, chatId)

	changePageMsg, err := b.generateChangePageMessage(chatId, page)
	if err != nil {
		return err
	}

	productsList = append(productsList, changePageMsg)

	b.sendMessages(productsList...)

	return nil
}

func (b *Bot) callbackAddToCart(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	cbSplit := strings.Split(callbackQuery.Data, " ")
	productId, err := strconv.Atoi(cbSplit[1])
	if err != nil {
		return err
	}

	order, err := b.services.GetOrder(chatId)
	if err != nil {
		return err
	}

	err = b.services.CreateCart(order.Id, productId)

	return err
}
