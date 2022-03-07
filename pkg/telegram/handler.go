package telegram

import (
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	commandMenu  = "menu"
)

const (
	callbackPreRegYes             = "pre_reg_yes"
	callbackRegLastYes            = "reg_last_yes"
	callbackRegLastNo             = "reg_last_no"
	callbackChangePage            = "ch_page"
	callbackAddCartKilo           = "add_cart_kilo"
	callbackAddCartBag            = "add_cart_bag"
	callbackReduceQuantity        = "reduce_quantity"
	callbackIncreaseQuantity      = "increase_quantity"
	callbackEditProfileYes        = "edit_profile_yes"
	callbackDeleteFromCart        = "delete_from_cart"
	callbackPlaceAnOrderYes       = "place_an_order_yes"
	callbackPlaceAnOrderNo        = "place_an_order_no"
	callbackEditOrderYes          = "edit_order_yes"
	callbackEditOrderNo           = "edit_order_no"
	callbackAcceptOrder           = "accept_order"
	callbackCourierRegLastYes     = "courier_reg_last_yes"
	callbackCourierRegLastNo      = "courier_reg_last_no"
	callbackEditCourierProfileYes = "edit_courier_profile_yes"
	callbackFinishOrder           = "finish_order"
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
		return nil
	}

	if strings.Contains(callbackQuery.Data, callbackAddCartKilo) {
		err := b.callbackAddToCart(callbackQuery)
		if err != nil {
			return err
		}
		return nil
	}

	if strings.Contains(callbackQuery.Data, callbackAddCartBag) {
		err := b.callbackAddToCart(callbackQuery)
		if err != nil {
			return err
		}
		return nil
	}

	if strings.Contains(callbackQuery.Data, callbackReduceQuantity) {
		err := b.callbackReduceQuantity(callbackQuery)
		if err != nil {
			return err
		}
		return nil
	}

	if strings.Contains(callbackQuery.Data, callbackIncreaseQuantity) {
		err := b.callbackIncreaseQuantity(callbackQuery)
		if err != nil {
			return err
		}
		return nil
	}

	if strings.Contains(callbackQuery.Data, callbackDeleteFromCart) {
		err := b.callbackDeleteFromCart(callbackQuery)
		if err != nil {
			return err
		}
		return nil
	}

	if strings.Contains(callbackQuery.Data, callbackAcceptOrder) {
		err := b.callbackAcceptOrder(callbackQuery)
		if err != nil {
			return err
		}
		return nil
	}

	if strings.Contains(callbackQuery.Data, callbackFinishOrder) {
		err := b.callbackFinishOrder(callbackQuery)
		if err != nil {
			return err
		}
		return nil
	}

	switch callbackQuery.Data {
	case callbackPreRegYes:
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
	case callbackCourierRegLastYes, callbackCourierRegLastNo:
		err := b.callbackCourierRegLast(callbackQuery)
		if err != nil {
			return err
		}
	case callbackEditCourierProfileYes:
		err := b.callbackEditCourierProfile(callbackQuery)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) handleStandartMessages(message *tgbotapi.Message) error {
	if message.Text == os.Getenv("COURIER_KEYWORD") {
		err := b.standartMessageCourierPreReg(message.Chat.ID)
		if err != nil {
			return err
		}
		return nil
	}

	user, err := b.services.GetUser(message.Chat.ID)
	if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set") {
		return err
	}

	if user.Id != 0 {
		err := b.registration(message)
		if err != nil {
			return err
		}

		err = b.placeAnOrder(message)
		if err != nil {
			return err
		}
		
		if user.DialogueStatus == "normal" {
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
			case "История заказов":
				err := b.standartMessageAllOrders(message.Chat.ID)
				if err != nil {
					return err
				}
			case "Профиль":
				err := b.standartMessageProfile(message.Chat.ID)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	courier, err := b.services.GetCourier(message.Chat.ID)
	if err != nil && err.Error() != "repository/GetCourier: select from couriers : sql: no rows in result set" {
		return err
	}
	if courier.Id != 0 {
		err := b.courierRegistration(message)
		if err != nil {
			return err
		}

		if courier.DialogueStatus == "normal" {
		switch message.Text {
			case "Активные заказы":
				err := b.standartMessageCourierActiveOrders(message.Chat.ID)
				if err != nil {
					return err
				}
			case "История заказов":
				err := b.standartMessageCourierAllOrders(message.Chat.ID)
				if err != nil {
					return err
				}
			case "Профиль":
				err := b.standartMessageCourierProfile(message.Chat.ID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
