package telegram

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//Регистрация
func (b *Bot) callbackPreReg(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	err := b.deleteReplyMarkup(callbackQuery.Message)
	if err != nil {
		return err
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "registration_add-name")
	if err != nil {
		return err
	}
	err = b.sendMessage(chatId, "Хорошо, укажите ваше имя")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) callbackRegLast(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID

	if callbackQuery.Data == callbackRegLastYes {
		err := b.sendMessageWithKeyboard(chatId, "Замечательно! Приятно с вами познакомиться :)", userMenuKeyboard)
		if err != nil {
			return err
		}
	}
	if callbackQuery.Data == callbackRegLastNo {
		err := b.services.UpdateUser(chatId, "dialogue_status", "registration_add-name")
		if err != nil {
			return err
		}

		err = b.sendMessage(chatId, "Хорошо. Пожалуйста, напишите ваше имя.")
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

func (b *Bot) callbackCourierRegLast(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID

	if callbackQuery.Data == callbackCourierRegLastYes {
		err := b.sendMessageWithKeyboard(chatId, "Замечательно! Приятно с вами познакомиться :)", courierMenuKeyboard)
		if err != nil {
			return err
		}
	}
	if callbackQuery.Data == callbackCourierRegLastNo {
		err := b.services.UpdateCourier(chatId, "dialogue_status", "registration_add-name")
		if err != nil {
			return err
		}

		err = b.sendMessage(chatId, "Хорошо. Пожалуйста, напишите ваше имя.")
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

//Профиль
func (b *Bot) callbackEditProfile(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	err := b.deleteReplyMarkup(callbackQuery.Message)
	if err != nil {
		return err
	}

	err = b.closeKeyboard(chatId, "В процессе редактирования профиля функционал бота не доступен")
	if err != nil {
		return err
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "registration_add-name")
	if err != nil {
		return err
	}

	err = b.sendMessage(chatId, "Хорошо. Пожалуйста, напишите ваше имя.")
	if err != nil {
		return err
	}

	return nil
}

//Каталог
func (b *Bot) callbackChangePage(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	cbSplit := strings.Split(callbackQuery.Data, " ")
	page, err := strconv.Atoi(cbSplit[1])
	if err != nil {
		return err
	}

	offset := (page - 1) * 5

	products, err := b.services.GetProducts(offset)
	if err != nil {
		return err
	}

	var prodOnPage int
	if (page - 2) * 5 < 0 {
		prodOnPage, err = b.services.CountProductsOnPage((page - 2) * -5)
		if err != nil {
			return err
		}
	} else {
		prodOnPage, err = b.services.CountProductsOnPage((page - 2) * -5)
		if err != nil {
			return err
		}
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

//Корзина
func (b *Bot) callbackAddToCart(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	cbSplit := strings.Split(callbackQuery.Data, " ")
	idPrice := strings.Split(cbSplit[1], "-")
	productId, err := strconv.Atoi(idPrice[0])
	if err != nil {
		return err
	}

	productPrice, err := strconv.Atoi(idPrice[1])
	if err != nil {
		return err
	}

	deliveryFormat := idPrice[2]

	order, err := b.services.GetOrderByUser(chatId)
	if err != nil {
		return err
	}

	err = b.services.CreateCart(order.Id, productId, productPrice, deliveryFormat)

	return err
}

func (b *Bot) callbackReduceQuantity(callbackQuery *tgbotapi.CallbackQuery) error {
	cbSplit := strings.Split(callbackQuery.Data, " ")
	productId, err := strconv.Atoi(cbSplit[1])
	if err != nil {
		return err
	}

	order, err := b.services.GetOrderByUser(callbackQuery.From.ID)
	if err != nil {
		return err
	}

	quantity, err := b.services.GetQuantity(order.Id, productId)
	if err != nil {
		return err
	}

	if cbSplit[0] == "reduce_quantity_10" {
		quantity -= 10
		err = b.services.UpdateQuantity(order.Id, productId, quantity)
		if err != nil {
			return err
		}
	}
	if cbSplit[0] == "reduce_quantity_1" {
		quantity -= 1
		err = b.services.UpdateQuantity(order.Id, productId, quantity)
		if err != nil {
			return err
		}
	}

	if quantity <= 0 {
		err = b.services.DeleteProductFromCart(order.Id, productId)
		if err != nil {
			return err
		}
		err = b.deleteMessage(callbackQuery.From.ID, callbackQuery.Message.MessageID)
		if err != nil {
			return err
		}
		return nil
	}

	newText, err := b.updatedCartProductCardText(order.Id, productId)
	if err != nil {
		return err
	}
	err = b.updateMessage(callbackQuery.From.ID, callbackQuery.Message.MessageID, callbackQuery.Message.ReplyMarkup, newText)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) callbackIncreaseQuantity(callbackQuery *tgbotapi.CallbackQuery) error {
	cbSplit := strings.Split(callbackQuery.Data, " ")
	productId, err := strconv.Atoi(cbSplit[1])
	if err != nil {
		return err
	}

	order, err := b.services.GetOrderByUser(callbackQuery.From.ID)
	if err != nil {
		return err
	}

	quantity, err := b.services.GetQuantity(order.Id, productId)
	if err != nil {
		return err
	}
	if cbSplit[0] == "increase_quantity_10" {
		quantity += 10
		err = b.services.UpdateQuantity(order.Id, productId, quantity)
		if err != nil {
			return err
		}
	}
	if cbSplit[0] == "increase_quantity_1" {
		quantity += 1
		err = b.services.UpdateQuantity(order.Id, productId, quantity)
		if err != nil {
			return err
		}
	}

	newText, err := b.updatedCartProductCardText(order.Id, productId)
	if err != nil {
		return err
	}
	err = b.updateMessage(callbackQuery.From.ID, callbackQuery.Message.MessageID, callbackQuery.Message.ReplyMarkup, newText)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) callbackDeleteFromCart(callbackQuery *tgbotapi.CallbackQuery) error {
	cbSplit := strings.Split(callbackQuery.Data, " ")
	productId, err := strconv.Atoi(cbSplit[1])
	if err != nil {
		return err
	}

	order, err := b.services.GetOrderByUser(callbackQuery.From.ID)
	if err != nil {
		return err
	}

	err = b.services.DeleteProductFromCart(order.Id, productId)
	if err != nil {
		return err
	}

	err = b.deleteMessage(callbackQuery.From.ID, callbackQuery.Message.MessageID)
	if err != nil {
		return err
	}
	return nil
}

//Оформление заказа
func (b *Bot) callbackPlaceAnOrderYes(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID

	err := b.deleteMessage(chatId, callbackQuery.Message.MessageID)
	if err != nil {
		return err
	}

	order, err := b.services.GetOrderByUser(chatId)
	if err != nil {
		return err
	}

	products, err := b.services.GetProductsFromCart(order.Id)
	if err != nil {
		return err
	}

	if len(products) == 0 {
		err = b.sendMessage(chatId, "Ваша корзина пуста")
		return err
	}

	b.closeKeyboard(chatId, "В процессе оформления заказа вы не можете использовать весь функционал бота")
	err = b.sendMessage(chatId, "Укажите адрес доставки.")
	if err != nil {
		return err
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "order_add-address")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) callbackPlaceAnOrderNo(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID

	err := b.deleteMessage(chatId, callbackQuery.Message.MessageID)
	if err != nil {
		return err
	}

	order, err := b.services.GetOrderByUser(chatId)
	if err != nil {
		return err
	}

	err = b.services.UpdateOrder(chatId, "order_status", "denied")
	if err != nil {
		return err
	}

	err = b.services.DeleteCart(order.Id)
	if err != nil {
		return err
	}

	err = b.sendMessage(chatId, "Ваш заказ отменен")
	if err != nil {
		return err
	}

	return nil
}

//Редактирование заказа
func (b *Bot) callbackEditOrderYes(callbackQuery *tgbotapi.CallbackQuery) error {
	err := b.deleteMessage(callbackQuery.From.ID, callbackQuery.Message.MessageID)
	if err != nil {
		return err
	}

	chatId := callbackQuery.From.ID

	err = b.services.UpdateUser(chatId, "dialogue_status", "order_add-name")
	if err != nil {
		return err
	}

	err = b.sendMessage(chatId, "Укажите имя заказчика.")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) callbackEditOrderNo(callbackQuery *tgbotapi.CallbackQuery) error {
	err := b.deleteMessage(callbackQuery.From.ID, callbackQuery.Message.MessageID)
	if err != nil {
		return err
	}

	chatId := callbackQuery.From.ID

	order, err := b.services.GetOrderByUser(chatId)
	if err != nil {
		return err
	}

	if order.UserName == "" || order.UserNumber == "" || order.UserCity == "" || order.DeliveryAdress == "" {
		err = b.services.UpdateUser(chatId, "dialogue_status", "order_add-name")
		if err != nil {
			return err
		}

		err = b.sendMessage(chatId, "Некорректные данные: Одно из полей пусто! Укажите имя.")
		if err != nil {
			return err
		}

		return nil
	}

	err = b.services.UpdateUser(chatId, "dialogue_status", "normal")
	if err != nil {
		return err
	}

	err = b.sendMessageToDeliveryService(chatId)
	if err != nil {
		return err
	}

	err = b.sendMessageWithKeyboard(chatId, fmt.Sprintf("Отлично! Ваш заказ №%d отправлен в службу доставки. Когда ваш заказ примут, вам придет сообщение с данными о курьере.", order.Id), userMenuKeyboard)
	if err != nil {
		return err
	}

	err = b.services.UpdateOrder(chatId, "order_status", "finished")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) updatedCartProductCardText(orderId int, productId int) (string, error) {
	cart, err := b.services.GetCart(orderId, productId)
	if err != nil {
		return "", err
	}

	product, err := b.services.GetProduct(cart.ProductId)
	if err != nil {
		return "", err
	}

	newText := fmt.Sprintf("%s x%d %s\n\nЦена: %d.%d₽\n\nОписание: %s", product.Title, cart.Quantity, cart.DeliveryFormat, cart.Price/100, cart.Price%100, product.Description)
	return newText, nil
}

func (b *Bot) callbackAcceptOrder(callbackQuery *tgbotapi.CallbackQuery) error {
	courier, err := b.services.GetCourier(callbackQuery.From.ID)
	if err != nil {
		return err
	}

	cbSplit := strings.Split(callbackQuery.Data, " ")
	orderId, err := strconv.Atoi(cbSplit[1])
	if err != nil {
		return err
	}

	user, err := b.services.GetOrderUser(orderId)
	if err != nil {
		return err
	}

	order, err := b.services.GetOrderById(orderId)
	if err != nil {
		return err
	}

	orderText, err := b.generateOrderText(order)
	if err != nil {
		return err
	}

	err = b.deleteReplyMarkup(callbackQuery.Message)
	if err != nil {
		return err
	}

	courierText := fmt.Sprintf("Ваш заказ №%d доставит:\n\t%s.\n\t%s.\n", orderId, courier.Name, courier.Number)

	err = b.services.CreateCourierOrder(orderId, courier.Id)
	if err != nil {
		return err
	}

	err = b.sendMessage(user.TelegramId, courierText)
	if err != nil {
		return err
	}

	err = b.sendMessageWithKeyboard(courier.TelegramId, orderText, tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Завершить заказ", fmt.Sprintf("finish_order %d", orderId)),
		),
	))
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) callbackEditCourierProfile(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	err := b.deleteReplyMarkup(callbackQuery.Message)
	if err != nil {
		return err
	}

	err = b.closeKeyboard(chatId, "В процессе редактирования профиля функционал бота не доступен")
	if err != nil {
		return err
	}

	err = b.services.UpdateCourier(chatId, "dialogue_status", "registration_add-name")
	if err != nil {
		return err
	}

	err = b.sendMessage(chatId, "Хорошо. Пожалуйста, напишите ваше имя.")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) callbackFinishOrder(callbackQuery *tgbotapi.CallbackQuery) error {
	cbSplit := strings.Split(callbackQuery.Data, " ")
	orderId, err := strconv.Atoi(cbSplit[1])
	if err != nil {
		return err
	}

	err = b.services.UpdateCourierOrder(orderId, "status", "delivered")
	if err != nil {
		return err
	}

	err = b.deleteReplyMarkup(callbackQuery.Message)
	if err != nil {
		return err
	}

	return nil
}
