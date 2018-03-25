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
	"set_key": setApiKey,
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
	return Message{ChatId: update.From.Id, Text: "API ключ успешно обновлен"}.Send()
}
