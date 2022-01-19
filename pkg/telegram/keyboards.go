package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var preRegistrationBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "pre_reg_yes"),
		tgbotapi.NewInlineKeyboardButtonData("🔴", "pre_reg_no"),
	),
)

var registrationLastBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "reg_last_yes"),
		tgbotapi.NewInlineKeyboardButtonData("🔴", "reg_last_no"),
	),
)