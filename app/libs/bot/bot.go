package bot

import (
	"errors"

	"fmt"

	"github.com/yarysh/car-online_bot/app"
)

type Bot struct {
	ApiKey string
}

func (b Bot) ProcessUpdate(update map[string]interface{}) (res string, err error) {
	message := b.getMessage(update)
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
	res, err := app.DbMap.Exec(
		"INSERT INTO users (username, first_name, last_name, chat_id, api_key) VALUES (:username, :firstName, :lastName, :chatId, :apiKey) "+
			"ON CONFLICT (chat_id) DO UPDATE SET "+
			"(username, first_name, last_name, chat_id, api_key) = (:username, :firstName, :lastName, :chatId, :apiKey)",
		map[string]interface{}{
			"username":  msg.From.Username,
			"firstName": msg.From.FirstName,
			"lastName":  msg.From.LastName,
			"chatId":    msg.From.Id,
			"apiKey":    payload,
		},
	)
	fmt.Println(res, err)
	return "", nil
}
