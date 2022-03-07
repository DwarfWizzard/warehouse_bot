package telegram

import (
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) sendMessage(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := b.bot.Send(msg)

	if err != nil {
		return fmt.Errorf("telegram/sendMessage: error %s", err.Error())
	}

	return err
}

func (b *Bot) sendMessages(args ...tgbotapi.Chattable) {
	for _, arg := range args {
		_, err := b.bot.Send(arg)
		if err != nil {
			b.services.PrintLog(fmt.Sprintf("telegram/sendMessages: %s", err.Error()), 1)
			continue
		}
	}
}

func (b *Bot) openKeyboard(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, "Клавиатура открыта")
	msg.ReplyMarkup = userMenuKeyboard

	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("telegram/openKeyboard: error %s", err.Error())
	}
	return err
}

func (b *Bot) closeKeyboard(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("telegram/closeKeyboard: error %s", err.Error())
	}
	return nil
}

func (b *Bot) sendMessageWithKeyboard(chatId int64, text string, Keyboard interface{}) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = Keyboard

	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("telegram/sendMessageWithKeyboard: error %s", err.Error())
	}
	return nil
}

func (b *Bot) updateMessage(chatId int64, messageId int, keyboard *tgbotapi.InlineKeyboardMarkup, newText string) error {
	editMsg := tgbotapi.NewEditMessageText(chatId, messageId, newText)
	editMsg.ReplyMarkup = keyboard
	_, err := b.bot.Send(editMsg)

	if err != nil {
		return fmt.Errorf("telegram/updateMessage: error %s", err.Error())
	}

	return nil
}

func (b *Bot) messageRegistrationLast(chatId int64) error {
	user, err := b.services.GetUser(chatId)
	if err != nil {
		return err
	}

	err = b.sendMessageWithKeyboard(chatId, fmt.Sprintf("Проверьте верность указанных данных. \nИмя: %s\nНомер: %s\nГород: %s", user.Name, user.Number, user.City), registrationLastBoard)
	return err
}

func (b *Bot)  messageCourierRegistrationLast(chatId int64) error {
	courier, err := b.services.GetCourier(chatId)
	if err != nil {
		return err
	}

	err = b.sendMessageWithKeyboard(chatId, fmt.Sprintf("Проверьте верность указанных данных. \nИмя: %s\nНомер: %s", courier.Name, courier.Number), courierLastBoard)
	return err
}

func (b *Bot) messageOrderLast(chatId int64) error {
	order, err := b.services.GetOrderByUser(chatId)
	if err != nil {
		return err
	}

	productListText, err := b.generateOrderText(order)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, productListText)
	_, err = b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("telegram/messageOrderLast: error %s", err.Error())
	}

	msg = tgbotapi.NewMessage(chatId, "Нажмите зеленую кнопку, чтобы завершить заказ.\nЖелтую, если хотите изменить данные о заказчике.\nКрасную, чтобы отменить заказ")
	msg.ReplyMarkup = editOrderUserInfoBoard
	_, err = b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("telegram/messageOrderLast: error %s", err.Error())
	}

	return nil
}

func (b *Bot) sendMessageToDeliveryService(chatId int64) error {
	order, err := b.services.GetOrderByUser(chatId)
	if err != nil {
		return err
	}

	productListText, err := b.generateOrderText(order)
	if err != nil {
		return err
	}
	
	subsidiary, err := b.services.GetSubsidiary(order.UserCity)
	if err != nil && strings.Contains(err.Error(), "error sql: no rows in result set") {
		err = b.services.UpdateUser(chatId, "dialogue_status", "order_add-city")
		if err != nil {
			return err
		}
	
		kb, err := b.CreateSubsidarysKB()
		if err != nil {
			return err
		}
	
		err = b.sendMessageWithKeyboard(chatId, "Данный филиал отсутсвует. Пожалуйста, выберите филиал из появившегося списка:", kb)
		if err != nil {
			return err
		}

		return errors.New("invalid subsidiary")
		
	} else if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(subsidiary.ChatId, productListText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Принять заказ", fmt.Sprintf("accept_order %d", order.Id)),
		),
	)
	_, err = b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("telegram/sendMessageToDeliveryService: error %s", err.Error())
	}

	return nil
}
