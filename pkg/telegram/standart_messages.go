package telegram

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"

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
	case "registration_add-city":
		err := b.standartMessageRegistrationCity(message)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) courierRegistration(message *tgbotapi.Message) error {
	courier, err := b.services.GetCourier(message.Chat.ID)
	if err != nil && err.Error() != ""{
		return err
	}

	switch courier.DialogueStatus {
	case "registration_add-name":
		err = b.courierRegistrationAddName(message)
		if err != nil {
			return err
		}
	case "registration_add-number":
		err = b.courierRegistrationAddNumber(message)
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
	case "order_add-city":
		err = b.standartMessageAddOrderCity(message)
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

	err = b.services.UpdateUser(chatId, "dialogue_status", "registration_add-city")
	if err != nil {
		return err
	}

	kb, err := b.CreateSubsidarysKB()
	if err != nil {
		return err
	}
	
	err = b.sendMessageWithKeyboard(chatId, "И теперь выберете город из списка.\nЕсли вашего города в списке нет, значит филлиал 'ПРОСТОРОС' в нем пока что отсутвует.", kb)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) standartMessageRegistrationCity(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateUser(chatId, "city", message.Text)
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

func (b *Bot) courierRegistrationAddName(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateCourier(chatId, "name", message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateCourier(chatId, "dialogue_status", "registration_add-number")
	if err != nil {
		return err
	}

	err = b.sendMessage(chatId, "Отлично, теперь укажите ваш номер телефона.")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) courierRegistrationAddNumber(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateCourier(chatId, "number", message.Text)
	if err != nil {
		return err
	}

	err = b.services.UpdateCourier(chatId, "dialogue_status", "normal")
	if err != nil {
		return err
	}

	err = b.messageCourierRegistrationLast(chatId)
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

	if checkAdress(message.Text) != nil {
		err := b.sendMessage(chatId, "Некорректный адрес! Укажите адрес заново.")
		if err != nil {
			return err
		}
		return nil
	}

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

func checkAdress(address string) error {
	msgSplit := strings.Split(address, " ")
	if len(msgSplit) < 2 {
		return errors.New("invalid address")
	}

	var count int = 0
	for _, r := range address {
		if unicode.IsDigit(r) {
			count++
		}
	}

	if count == 0 {
		return errors.New("invalid address")
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

	err = b.services.UpdateUser(chatId, "dialogue_status", "order_add-city")
	if err != nil {
		return err
	}

	kb, err := b.CreateSubsidarysKB()
	if err != nil {
		log.Println(1111111)
		return err
	}

	err = b.sendMessageWithKeyboard(chatId, "Выберите филиал из появившегося списка", kb)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) standartMessageAddOrderCity(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	err := b.services.UpdateOrder(chatId, "user_city", message.Text)
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
	order, err := b.services.GetOrderByUser(chatId)
	if err != nil {
		return err
	}

	products, err := b.services.GetProductsFromCart(order.Id)
	if err != nil {
		return err
	}

	carts, err := b.services.GetCarts(order.Id)
	if err != nil {
		return err
	}

	productsList := b.generateShopingCartProductCards(products, carts,chatId)

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

	b.sendMessageWithKeyboard(chatId, fmt.Sprintf("Ваш профиль.\n\nИмя: %s\n\nНомер телефона: %s\n\nГород: %s\n\nРедактировать профиль?", user.Name, user.Number, user.City), editProfileBoard)
	return nil
}

func (b *Bot) standartMessageCourierPreReg(chatId int64) error{
	courier, err := b.services.GetCourier(chatId)
	if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set"){
		return err
	}
	
	err = b.services.DeleteUser(chatId)
	if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set"){
		return err
	}

	if courier.Id == 0 {
		err := b.services.CreateCourier(chatId)
		if err != nil {
			return err
		}

		err = b.services.UpdateCourier(chatId, "dialogue_status", "registration_add-name")
		if err != nil {
			return err
		}

		b.sendMessage(chatId, "Здравствуйте, курьер!\nПожалуйста, напишите ваше имя.")
	} else {
		b.sendMessage(chatId, "Вы уже зарегистрированы в системе.")
	}
	return nil
}

func (b *Bot) standartMessageAllOrders(chatId int64) error {
	user, err := b.services.GetUser(chatId)
	if err != nil {
		return err
	}

	orders, err := b.services.GetAllOrdersUser(user.Id)
	if err != nil {
		return err
	}

	orderCards := b.generateAllOrderCards(orders, chatId)
	b.sendMessages(orderCards...)

	return nil
}

func (b *Bot) standartMessageCourierActiveOrders(chatId int64) error {
	courier, err := b.services.GetCourier(chatId)
	if err != nil {
		return err
	}

	orders, err := b.services.GetActiveOrders(courier.Id)
	if err != nil {
		return err
	}

	orderCards := b.generateActiveOrderCards(orders, chatId)
	
	b.sendMessages(orderCards...)

	return nil
}

func (b *Bot) standartMessageCourierAllOrders(chatId int64) error {
	courier, err := b.services.GetCourier(chatId)
	if err != nil {
		return err
	}

	orders, err := b.services.GetCourierOrders(courier.Id)
	if err != nil {
		return err
	}

	orderCards := b.generateAllOrderCards(orders, chatId)
	
	b.sendMessages(orderCards...)

	return nil
}

func (b *Bot) standartMessageCourierProfile(chatId int64) error {
	courier, err := b.services.GetCourier(chatId)
	if err != nil {
		return err
	}
	
	err = b.sendMessageWithKeyboard(chatId, fmt.Sprintf("Ваш профиль.\n\nИмя: %s\n\nНомер телефона: %s\n\nРедактировать профиль?", courier.Name, courier.Number), editCourierProfileBoard)
	if err != nil {
		return err
	}

	return nil
}