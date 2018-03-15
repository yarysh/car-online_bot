package bot

import (
	"errors"
	"strconv"

	"github.com/yarysh/car-online_bot/app/libs/helpers"
	"github.com/yarysh/car-online_bot/app/models"
)

func ProcessUpdate(update map[string]interface{}) (map[string]interface{}, error) {
	message := getMessage(update)
	name, payload := message.getCommandInfo()
	if name == "" {
		return nil, nil
	}
	command, ok := commands[name]
	if !ok {
		return nil, errors.New("no such commands")
	}
	return command(payload, message)
}

func SendMessage(chatId int64, text string) (map[string]interface{}, error) {
	return helpers.GetJsonResponse(getApiUrl("sendMessage", map[string]string{
		"chat_id": strconv.Itoa(int(chatId)), "text": text,
	}))
}

// Map of command name and func executor
var commands = map[string]func(payload string, message message) (map[string]interface{}, error){
	"set_key": setApiKey,
}

// Set car-online api key for user
func setApiKey(payload string, msg message) (map[string]interface{}, error) {
	if payload == "" {
		return nil, errors.New("api key is empty")
	}
	_, err := models.User{
		Username:  msg.From.Username,
		FirstName: msg.From.FirstName,
		LastName:  msg.From.LastName,
		ChatId:    msg.From.Id,
		ApiKey:    payload,
	}.UpdateApiKey()
	if err != nil {
		return nil, errors.New("sql error")
	}
	return SendMessage(msg.From.Id, "API ключ успешно обновлен")
}
