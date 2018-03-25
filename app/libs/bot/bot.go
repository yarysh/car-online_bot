package bot

import (
	"errors"

	"github.com/yarysh/car-online_bot/app/models"
)

func ProcessUpdate(update map[string]interface{}) (map[string]interface{}, error) {
	updateMessage := getUpdateMessage(update)
	name, payload := updateMessage.getCommandInfo()
	if name == "" {
		return nil, nil
	}
	command, ok := commands[name]
	if !ok {
		return nil, errors.New("no such commands")
	}
	return command(payload, updateMessage)
}

// Map of command name and func executor
var commands = map[string]func(payload string, update updateMessage) (map[string]interface{}, error){
	"start":   start,
	"set_key": setApiKey,
}

func start(payload string, update updateMessage) (map[string]interface{}, error) {
	m := Message{
		ChatId: update.From.Id,
		Text: "Этот бот позволяет получить информацию об автомобиле с сайта [car-online.ru](http://car-online.ru/)\n" +
			"Для того, чтобы начать пользоваться ботом, необходимо ввести ключ для API Car-Online, который вы можете найти в [ЛК](https://panel.car-online.ru)\n" +
			"Для доступа ко всем функциям бота вы также можете использовать тестовый ключ: *test*.",
		ParseMode: "Markdown",
		ReplyMarkup: InlineKeyboard{InlineKeyboard: [][]InlineKeyboardButton{
			{InlineKeyboardButton{Text: "Установить ключ", CallbackData: "/set_key"}},
			{InlineKeyboardButton{Text: "Установить тестовый ключ", CallbackData: "/set_key test"}},
			{InlineKeyboardButton{Text: "Помощь", CallbackData: "/help"}},
		}},
	}
	return m.Send()
}

// Set car-online api key for user
func setApiKey(payload string, update updateMessage) (map[string]interface{}, error) {
	if payload == "" {
		return nil, errors.New("api key is empty")
	}
	_, err := models.User{
		Username:  update.From.Username,
		FirstName: update.From.FirstName,
		LastName:  update.From.LastName,
		ChatId:    update.From.Id,
		ApiKey:    payload,
	}.UpdateApiKey()
	if err != nil {
		return nil, errors.New("sql error")
	}
	m := Message{ChatId: update.From.Id, Text: "API ключ успешно обновлен"}
	return m.Send()
}
