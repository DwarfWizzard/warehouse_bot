package telegram

import (
	"errors"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) generateProductsCardMessages(chatId int64, page int) ([]tgbotapi.MessageConfig, error) {
	var productsCards []tgbotapi.MessageConfig
	productsNum, err := b.services.CountAllProducts()
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * 5

	products, err := b.services.GetProducts(offset)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		productText := fmt.Sprintf("%s\n\nЦена:%s\n\nОписание:%s", product.Title, product.Price, product.Description)
		productCard := tgbotapi.NewMessage(chatId, productText)
		productCard.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", fmt.Sprintf("add_cart %d", product.Id)),
			),
		)

		productsCards = append(productsCards, productCard)
	}
	pageMsg := tgbotapi.NewMessage(chatId, fmt.Sprintf("%d/%d", page, productsNum/5))
	var pageChangeKeyboard tgbotapi.InlineKeyboardMarkup

	if page == 1 {
		pageChangeKeyboard, err = b.generateInlineKeyBoard(1, map[string]string{
			">": fmt.Sprintf("ch_page %d", page+1),
		})
		if err != nil {
			return productsCards, err
		}
	} else if page == productsNum/5 {
		pageChangeKeyboard, err = b.generateInlineKeyBoard(1, map[string]string{
			"<": fmt.Sprintf("ch_page %d", page-1),
		})
		if err != nil {
			return productsCards, err
		}
	} else {
		pageChangeKeyboard, err = b.generateInlineKeyBoard(2, map[string]string{
			"<": fmt.Sprintf("ch_page %d", page-1),
			">": fmt.Sprintf("ch_page %d", page+1),
		})
		if err != nil {
			return productsCards, err
		}
	}

	pageMsg.ReplyMarkup = pageChangeKeyboard
	productsCards = append(productsCards, pageMsg)
	return productsCards, err
}

func (b *Bot) generateInlineKeyBoard(buttonNum int, textAndData map[string]string) (tgbotapi.InlineKeyboardMarkup, error) {
	var keyboard tgbotapi.InlineKeyboardMarkup
	buttons := make([]tgbotapi.InlineKeyboardButton, buttonNum)

	if len(textAndData) != buttonNum {
		return keyboard, errors.New("number of text and data for keyboadrd not equal number of buttons")
	}

	var buttonTexts []string
	for t := range textAndData {
		buttonTexts = append(buttonTexts, t)
	}

	var buttonData []string
	for _, d := range textAndData {
		buttonData = append(buttonData, d)
	}

	for i := 0; i < buttonNum; i++ {
		buttons[i] = tgbotapi.NewInlineKeyboardButtonData(buttonTexts[i], buttonData[i])
	}

	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			buttons...
		),
	)

	return keyboard, nil
}