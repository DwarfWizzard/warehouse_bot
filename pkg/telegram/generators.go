package telegram

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) generateProductListCardsMessages(products []models.Product, chatId int64) []tgbotapi.Chattable {
	var productsCards []tgbotapi.Chattable
	for _, product := range products {
		productText := fmt.Sprintf("%s\n\nРозничная цена: %d.%d₽\nОптовая цена: %d.%d₽\n\nОписание: %s", product.Title, product.PriceKilo/100, product.PriceKilo%100, product.PriceBag/100, product.PriceBag%100, product.Description)

		photoBytes, err := ioutil.ReadFile(os.Getenv("IMAGE_PATH") + "/" + product.ImageName)
		if err != nil {
			productCard := tgbotapi.NewMessage(chatId, productText)
			productCard.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину кг", fmt.Sprintf("add_cart_kilo %d-%d-%s", product.Id, product.PriceKilo, "кг")),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину мешок", fmt.Sprintf("add_cart_bag %d-%d-%s", product.Id, product.PriceBag, "мешок")),
				),
			)
			productsCards = append(productsCards, productCard)
		} else {
			log.Println(product.ImageName)
			file := tgbotapi.FileBytes{
				Name:  product.ImageName,
				Bytes: photoBytes,
			}

			productCard := tgbotapi.NewPhoto(chatId, file)
			productCard.Caption = productText
			productCard.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину кг", fmt.Sprintf("add_cart_kilo %d-%d-%s", product.Id, product.PriceKilo, "кг")),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину мешок", fmt.Sprintf("add_cart_bag %d-%d-%s", product.Id, product.PriceBag, "мешок")),
				),
			)
			productsCards = append(productsCards, productCard)
		}

	}
	return productsCards
}

func (b *Bot) generateShopingCartProductCards(products []models.Product, carts []models.ShopingCart, chatId int64) []tgbotapi.Chattable {
	var productsCards []tgbotapi.Chattable
	for i, product := range products {
		productText := fmt.Sprintf("%s x%d %s\n\nЦена: %d.%d₽\n\nОписание: %s", product.Title, carts[i].Quantity, carts[i].DeliveryFormat, carts[i].Price/100, carts[i].Price%100, product.Description)
		productCard := tgbotapi.NewMessage(chatId, productText)
		productCard.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("-10", fmt.Sprintf("reduce_quantity_10 %d", product.Id)),
				tgbotapi.NewInlineKeyboardButtonData("-1", fmt.Sprintf("reduce_quantity_1 %d", product.Id)),
				tgbotapi.NewInlineKeyboardButtonData("+1", fmt.Sprintf("increase_quantity_1 %d", product.Id)),
				tgbotapi.NewInlineKeyboardButtonData("+10", fmt.Sprintf("increase_quantity_10 %d", product.Id)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Убрать товар", fmt.Sprintf("delete_from_cart %d", product.Id)),
			),
		)

		productsCards = append(productsCards, productCard)
	}
	return productsCards
}

func (b *Bot) generateChangePageMessage(chatId int64, page int) (tgbotapi.MessageConfig, error) {
	productsNum, err := b.services.CountProducts()
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}

	var pageNum int
	if productsNum%5 != 0 {
		pageNum = (productsNum / 5) + 1
	} else {
		pageNum = productsNum / 5
	}

	pageMsg := tgbotapi.NewMessage(chatId, fmt.Sprintf("%d/%d", page, pageNum))

	if productsNum > 5 {
		var state int
		if page == 1 {
			state = 1
		} else if page == pageNum {
			state = 2
		} else {
			state = 3
		}

		pageMsg.ReplyMarkup = generateChangeKeyboard(state, fmt.Sprintf("ch_page %d", page-1), fmt.Sprintf("ch_page %d", page+1))
	}

	return pageMsg, nil
}

func generateChangeKeyboard(state int, valueLeft string, valueRight string) tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	buttonLeft := tgbotapi.NewInlineKeyboardButtonData("<", valueLeft)
	buttonRight := tgbotapi.NewInlineKeyboardButtonData(">", valueRight)
	if state == 1 {
		buttons = append(buttons, buttonRight)
	}
	if state == 2 {
		buttons = append(buttons, buttonLeft)
	}
	if state == 3 {
		buttons = append(buttons, buttonLeft, buttonRight)
	}

	productsChangeKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			buttons...,
		),
	)
	return productsChangeKeyboard
}

func (b *Bot) generateActiveOrderCards(orders []models.Order, chatId int64) []tgbotapi.Chattable {
	var orderCards []tgbotapi.Chattable
	for _, order := range orders {
		orderText, err := b.generateOrderText(order)
		if err != nil {
			log.Println(err)
			continue
		}
		orderCard := tgbotapi.NewMessage(chatId, orderText)
		orderCard.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Завершить заказ", fmt.Sprintf("finish_order %d", order.Id)),
			),
		)

		orderCards = append(orderCards, orderCard)
	}

	return orderCards
}

func (b *Bot) generateAllOrderCards(orders []models.Order, chatId int64) []tgbotapi.Chattable {
	var orderCards []tgbotapi.Chattable
	for _, order := range orders {
		orderText, err := b.generateOrderText(order)
		if err != nil {
			b.services.PrintLog(err.Error(), 1)
			continue
		}

		orderCard := tgbotapi.NewMessage(chatId, orderText)

		orderCards = append(orderCards, orderCard)
	}

	return orderCards
}

func (b *Bot) generateOrderText(order models.Order) (string, error) {
	products, err := b.services.GetProductsFromCart(order.Id)
	if err != nil {
		return "", err
	}

	carts, err := b.services.GetCarts(order.Id)
	if err != nil {
		return "", err
	}

	var productListText string = fmt.Sprintf("Заказ №%d\n\nСписок продуктов:\n", order.Id)
	var totalPrice int
	for i, product := range products {
		productListText += fmt.Sprintf("\t%d. %s x%d %s %d.%d₽\n\n", i+1, product.Title, carts[i].Quantity, carts[i].DeliveryFormat, (carts[i].Price)/100, (carts[i].Price)%100)
		totalPrice += carts[i].Price
	}

	productListText += fmt.Sprintf("Общая сумма заказа: %d.%d₽\n\nДанные о заказчике:\n\tИмя заказчика: %s\n\tНомер заказчика: %s\n\tАдрес доставки: %s", totalPrice/100, totalPrice%100, order.UserName, order.UserNumber, order.DeliveryAdress)

	return productListText, nil
}
