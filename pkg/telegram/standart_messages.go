package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) registration(message *tgbotapi.Message) error {
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

func (b *Bot) placeAnOrder(message *tgbotapi.Message) error {
	user, err := b.services.GetUser(message.Chat.ID)
	if err != nil {
		return err
	}

	switch user.DialogueStatus {
	case "order_add-address":
		err = b.standartMessageAddOrderAddress(message)
		if err != nil {
			return err
		}

	case "order_add-name":
		err = b.standartMessageAddOrderName(message)
		if err != nil {
			return err
		}
	case "order_add-number":
		err = b.standartMessageAddOrderNumber(message)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) standartMessageRegistrationName(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateUser(chatId, "name", message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "registration_add-number")
	if err != nil {
		return err
	}

	err = b.sendMessage(chatId, "Отлично, теперь укажите ваш номер телефона.")
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) standartMessageRegistrationNumber(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateUser(chatId, "number", message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "normal")
	if err != nil {
		return err
	}

	err = b.messageRegistrationLast(chatId)
	if err != nil {
		return err
	}
	return nil
}
func (b *Bot) standartMessageCatalog(chatId int64) error {
	products, err := b.services.GetProducts(0)
	if err != nil {
		return err
	}

	productsList := b.generateProductListCardsMessages(products, chatId)

	changePageMsg, err := b.generateChangePageMessage(chatId, 1)
	if err != nil {
		return err
	}

	productsList = append(productsList, changePageMsg)
	b.sendMessages(productsList...)

	return nil
}

func (b *Bot) standartMessageAddOrderAddress(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateOrder(chatId, "delivery_adress", message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "normal")
	if err != nil {
		return err
	}

	err = b.messageOrderLast(chatId)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) standartMessageAddOrderName(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateOrder(chatId, "user_name", message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "order_add-number")
	if err != nil {
		return err
	}

	err = b.sendMessage(chatId, "Укажите номер телефона заказчика.")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) standartMessageAddOrderNumber(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateOrder(chatId, "user_number", message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "order_add-address")
	if err != nil {
		return err
	}

	err = b.sendMessage(chatId, "Укажите адрес доставки")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) standartMessageShopingCart(chatId int64) error {
	order, err := b.services.GetOrder(chatId)
	if err != nil {
		return err
	}

	products, err := b.services.GetProductsFromCart(order.Id)
	if err != nil {
		return err
	}

	productsList := b.generateShopingCartProductCards(products, chatId)

	b.sendMessages(productsList...)
	if len(products) == 0 {
		err = b.sendMessage(chatId, "Ваша корзина пуста")
		return err
	}
	err = b.sendMessageWithKeyboard(chatId, "Нажмите на зеленую кнопку, чтобы оформить заказ.\nКрасную, чтобы отказаться.", placeAnOrderBoard)

	return err
}

func (b *Bot) standartMessageProfile(chatId int64) error {
	user, err := b.services.GetUser(chatId)
	if err != nil {
		return err
	}

	b.sendMessageWithKeyboard(chatId, fmt.Sprintf("Ваш профиль.\n\nИмя: %s\n\nНомер телефона: %s\n\nРедактировать профиль?", user.Name, user.Number), editProfileBoard)
	return nil
}
