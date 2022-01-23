package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) sendMessage(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) sendMessages(args ...tgbotapi.MessageConfig) {
	for _, arg := range args {
		_, err := b.bot.Send(arg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func (b *Bot) openKeyboard(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, "Клавиатура открыта")
	msg.ReplyMarkup = menuKeyboard

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) closeKeyboard(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) sendMessageWithKeyboard(chatId int64, text string, Keyboard interface{}) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = Keyboard

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) updateMessage(chatId int64, messageId int, keyboard *tgbotapi.InlineKeyboardMarkup, newText string) error {
	editMsg := tgbotapi.NewEditMessageText(chatId, messageId, newText)
	editMsg.ReplyMarkup = keyboard
	_, err := b.bot.Send(editMsg)

	return err
}

func (b *Bot) messageRegistrationLast(chatId int64) error {
	user, err := b.services.GetUser(chatId)
	if err != nil {
		return err
	}

	err = b.sendMessageWithKeyboard(chatId, fmt.Sprintf("Проверьте верность указанных данных. \nИмя: %s\nНомер: %s", user.Name, user.Number), registrationLastBoard)
	return err
}

func (b *Bot) messageOrderLast(chatId int64) error {
	order, err := b.services.GetOrder(chatId)
	if err != nil {
		return err
	}

	products, err := b.services.GetProductsFromCart(order.Id)
	if err != nil {
		return err
	}

	var productListText string = "Список продуктов:\n"
	var totalPrice int
	for i, product := range products {
		productListText += fmt.Sprintf("\t%d. %s x%d %d.%d₽\n\n", i+1, product.Title, product.Quantity, (product.Price*product.Quantity)/100, (product.Price*product.Quantity)%100)
		totalPrice += product.Price * product.Quantity
	}

	productListText += fmt.Sprintf("Общая сумма заказа: %d.%d₽\n\nДанные о заказчике:\n\tИмя заказчика: %s\n\tНомер заказчика: %s\n\tАдрес доставки: %s", totalPrice/100, totalPrice%100, order.UserName, order.UserNumber, order.DeliveryAdress)

	msg := tgbotapi.NewMessage(chatId, productListText)
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	msg = tgbotapi.NewMessage(chatId, "Изменить личные данные для заказа?")
	msg.ReplyMarkup = editOrderUserInfoBoard
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
