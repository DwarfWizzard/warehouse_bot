package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart    = "start"
	commandMenu = "menu"
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
	callbackEditProfileYes   = "edit_profile_yes"
	callbackDeleteFromCart   = "delete_from_cart"
	callbackPlaceAnOrderYes  = "place_an_order_yes"
	callbackPlaceAnOrderNo   = "place_an_order_no"
	callbackEditOrderYes     = "edit_order_yes"
	callbackEditOrderNo     = "edit_order_no"
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.commandStart(message)
	case commandMenu:
		return b.openKeyboard(message.Chat.ID)
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

	if strings.Contains(callbackQuery.Data, callbackDeleteFromCart) {
		err := b.callbackDeleteFromCart(callbackQuery)
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
	case callbackEditProfileYes:
		err := b.callbackEditProfile(callbackQuery)
		if err != nil {
			return err
		}
	case callbackPlaceAnOrderYes:
		err := b.callbackPlaceAnOrderYes(callbackQuery)
		if err != nil {
			return err
		}
	case callbackPlaceAnOrderNo:
		err := b.callbackPlaceAnOrderNo(callbackQuery)
		if err != nil {
			return err
		}
	case callbackEditOrderYes:
		err := b.callbackEditOrderYes(callbackQuery)
		if err != nil {
			return err
		}
	case callbackEditOrderNo:
		err := b.callbackEditOrderNo(callbackQuery)
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

	err = b.placeAnOrder(message)
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
	case "Профиль":
		err := b.standartMessageProfile(message.Chat.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
