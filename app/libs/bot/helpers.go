package bot

import (
	"net/url"
	"regexp"

	"encoding/json"

	"net/http"

	"bytes"

	"io/ioutil"

	"fmt"

	"github.com/revel/revel"
	"github.com/yarysh/car-online_bot/app/libs/helpers"
)

func (m *Message) Send() (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonM, err := json.Marshal(m)
	fmt.Println(fmt.Sprintf("%s", jsonM))
	if err != nil {
		return result, err
	}
	resp, err := http.Post(getApiUrl("sendMessage", nil), "application/json", bytes.NewBuffer(jsonM))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return result, jsonErr
	}
	return result, nil
}

func getApiUrl(methodName string, params map[string]string) string {
	p := url.Values{}
	for key, value := range params {
		p.Add(key, value)
	}
	return "https://api.telegram.org/bot" + revel.Config.StringDefault("bot.api_key", "") + "/" + methodName + "?" + p.Encode()
}

func getUpdateMessage(update map[string]interface{}) updateMessage {
	result := updateMessage{}

	msg, ok := update["message"].(map[string]interface{})
	if !ok {
		return result
	}

	result.Id = int64(helpers.FloatDefault(msg, "message_id", 0))
	frm, ok := msg["from"].(map[string]interface{})
	if ok {
		result.From = from{
			Id:           int64(helpers.FloatDefault(frm, "id", 0)),
			IsBot:        helpers.BoolDefault(frm, "is_bot", false),
			FirstName:    helpers.StringDefault(frm, "first_name", ""),
			LastName:     helpers.StringDefault(frm, "last_name", ""),
			Username:     helpers.StringDefault(frm, "username", ""),
			LanguageCode: helpers.StringDefault(frm, "language_code", ""),
		}
	}
	cht, ok := msg["chat"].(map[string]interface{})
	if ok {
		result.Chat = chat{
			Id:        int64(helpers.FloatDefault(cht, "id", 0)),
			Type:      helpers.StringDefault(cht, "type", ""),
			FirstName: helpers.StringDefault(cht, "first_name", ""),
			LastName:  helpers.StringDefault(cht, "last_name", ""),
			Username:  helpers.StringDefault(cht, "username", ""),
		}
	}
	result.Text = helpers.StringDefault(msg, "text", "")
	entities, ok := msg["entities"].([]interface{})
	if ok {
		for _, entity := range entities {
			result.Entities = append(result.Entities, messageEntity{
				Type:   helpers.StringDefault(entity.(map[string]interface{}), "type", ""),
				Offset: int(helpers.FloatDefault(entity.(map[string]interface{}), "offset", 0)),
				Length: int(helpers.FloatDefault(entity.(map[string]interface{}), "length", 0)),
				Url:    helpers.StringDefault(entity.(map[string]interface{}), "url", ""),
			})
		}
	}
	return result
}

func (u updateMessage) getCommandInfo() (name string, payload string) {
	if len(u.Entities) == 0 || u.Entities[0].Type != "bot_command" {
		return "", ""
	}
	found := regexp.MustCompile(`/(?P<Command>\w+)\s(?P<Payload>\w+)`).FindStringSubmatch(u.Text)
	if found == nil || len(found) < 2 {
		return "", ""
	}
	if len(found) >= 3 {
		return found[1], found[2]
	} else {
		return found[1], ""
	}
}
