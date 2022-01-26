package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var preRegistrationBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "pre_reg_yes"),
	),
)

var registrationLastBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "reg_last_yes"),
		tgbotapi.NewInlineKeyboardButtonData("🔴", "reg_last_no"),
	),
)

var courierLastBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "courier_reg_last_yes"),
		tgbotapi.NewInlineKeyboardButtonData("🔴", "courier_reg_last_no"),
	),
)

var editProfileBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "edit_profile_yes"),
	),
)

var editOrderUserInfoBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "edit_order_yes"),
		tgbotapi.NewInlineKeyboardButtonData("🔴", "edit_order_no"),
	),
)

var editCourierProfileBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "edit_courier_profile_yes"),
	),
)

var placeAnOrderBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🟢", "place_an_order_yes"),
		tgbotapi.NewInlineKeyboardButtonData("🔴", "place_an_order_no"),
	),
)

var userMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Каталог"),
		tgbotapi.NewKeyboardButton("Корзина"),
		tgbotapi.NewKeyboardButton("Профиль"),
	),
)

var courierMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Активные заказы"),
		tgbotapi.NewKeyboardButton("История заказов"),
		tgbotapi.NewKeyboardButton("Профиль"),
	),
)
