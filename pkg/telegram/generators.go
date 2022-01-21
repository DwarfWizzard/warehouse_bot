package telegram

import (
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) generateProductListCardsMessages(products []models.Product, chatId int64) []tgbotapi.MessageConfig {
	var productsCards []tgbotapi.MessageConfig
	for _, product := range products {
		productText := fmt.Sprintf("%s\n\nЦена:%s₽\n\nОписание:%s", product.Title, product.Price, product.Description)
		productCard := tgbotapi.NewMessage(chatId, productText)
		productCard.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", fmt.Sprintf("add_cart %d", product.Id)),
			),
		)

		productsCards = append(productsCards, productCard)
	}
	return productsCards
}

func (b *Bot) generateShopingCartProductCards(products []models.Product, chatId int64) []tgbotapi.MessageConfig {
	var productsCards []tgbotapi.MessageConfig
	for _, product := range products {
		productText := fmt.Sprintf("%s\n\nЦена:%s\n\nОписание:%s", product.Title, product.Price, product.Description)
		productCard := tgbotapi.NewMessage(chatId, productText)
		productCard.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("-10", fmt.Sprintf("reduce_quantity_10 %d", product.Id)),
				tgbotapi.NewInlineKeyboardButtonData("-1", fmt.Sprintf("reduce_quantity_1 %d", product.Id)),
				tgbotapi.NewInlineKeyboardButtonData("+1", fmt.Sprintf("increase_quantity_10 %d", product.Id)),
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
		return  tgbotapi.MessageConfig{}, err
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
			buttons...
		),
	)
	return productsChangeKeyboard
}