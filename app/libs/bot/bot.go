package bot

import (
	"errors"

	"github.com/yarysh/car-online_bot/app/models"
)

func ProcessUpdate(update map[string]interface{}) (res string, err error) {
	message := getMessage(update)
	name, payload := message.getCommandInfo()
	if name == "" {
		return "", nil
	}
	command, ok := commands[name]
	if !ok {
		return "", errors.New("no such commands")
	}
	return command(payload, message)
}

// Map of command name and func executor
var commands = map[string]func(payload string, message message) (res string, err error){
	"set_key": setApiKey,
}

// Set car-online api key for user
func setApiKey(payload string, msg message) (string, error) {
	if payload == "" {
		return "", errors.New("api key is empty")
	}
	_, err := models.User{
		Username:  msg.From.Username,
		FirstName: msg.From.FirstName,
		LastName:  msg.From.LastName,
		ChatId:    msg.From.Id,
		ApiKey:    payload,
	}.UpdateApiKey()
	if err != nil {
		return "", errors.New("sql error")
	}
	return "API ключ успешно обновлен", nil
}
