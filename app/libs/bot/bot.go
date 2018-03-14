package bot

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"io/ioutil"

	"encoding/json"

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
	var result map[string]interface{}
	params := url.Values{}
	params.Add("chat_id", strconv.Itoa(int(chatId)))
	params.Add("text", text)
	resp, err := http.Get(getApiUrl("sendMessage") + "?" + params.Encode())
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return result, jsonErr
	}
	return result, err
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
