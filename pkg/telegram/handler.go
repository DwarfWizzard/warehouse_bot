package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart    = "start"
	commandProdList = "products"
)

const (
	callbackPreRegYes        = "pre_reg_yes"
	callbackPreRegNo         = "pre_reg_no"
	callbackRegLastYes       = "reg_last_yes"
	callbackRegLastNo        = "reg_last_no"
	callbackChangePage       = "ch_page"
	callbackAddCart          = "add_cart"
	callbackReduceQuantity   = "reduce_quantity"
	callbackIncreaseQuantity = "increase_quantity"
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

	if strings.Contains(callbackQuery.Data, callbackChangePage) {
		err := b.callbackChangePage(callbackQuery)
		if err != nil {
			return err
		}
	}

	if strings.Contains(callbackQuery.Data, callbackAddCart) {
		err := b.callbackAddToCart(callbackQuery)
		if err != nil {
			return err
		}
	}

	if strings.Contains(callbackQuery.Data, callbackReduceQuantity) {
		err := b.callbackReduceQuantity(callbackQuery)
		if err != nil {
			return err
		}
	}

	if strings.Contains(callbackQuery.Data, callbackIncreaseQuantity) {
		err := b.callbackIncreaseQuantity(callbackQuery)
		if err != nil {
			return err
		}
	}

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
	err := b.registration(message)
	if err != nil {
		return err
	}

	switch message.Text {
	case "Каталог":
		err := b.standartMessageCatalog(message.Chat.ID)
		if err != nil {
			return err
		}
	case "Корзина":
		err := b.standartMessageShopingCart(message.Chat.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
