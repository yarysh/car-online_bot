package bot

import (
	"net/url"
	"regexp"

	"github.com/revel/revel"
	"github.com/yarysh/car-online_bot/app/libs/helpers"
)

func getApiUrl(methodName string, params map[string]string) string {
	p := url.Values{}
	for key, value := range params {
		p.Add(key, value)
	}
	return "https://api.telegram.org/bot" + revel.Config.StringDefault("bot.api_key", "") + "/" + methodName + "?" + p.Encode()
}

func getMessage(update map[string]interface{}) message {
	result := message{}

	msg, ok := update["message"].(map[string]interface{})
	if !ok {
		return result
	}

	result.Id = int64(helpers.GetValue(msg, "message_id", 0).(float64))
	from, ok := msg["from"].(map[string]interface{})
	if ok {
		result.From = user{
			Id:           int64(helpers.GetValue(from, "id", 0).(float64)),
			IsBot:        helpers.GetValue(from, "is_bot", false).(bool),
			FirstName:    helpers.GetValue(from, "first_name", "").(string),
			LastName:     helpers.GetValue(from, "last_name", "").(string),
			Username:     helpers.GetValue(from, "username", "").(string),
			LanguageCode: helpers.GetValue(from, "language_code", "").(string),
		}
	}
	cht, ok := msg["chat"].(map[string]interface{})
	if ok {
		result.Chat = chat{
			Id:        int64(helpers.GetValue(cht, "id", 0).(float64)),
			Type:      helpers.GetValue(cht, "type", "").(string),
			FirstName: helpers.GetValue(cht, "first_name", "").(string),
			LastName:  helpers.GetValue(cht, "last_name", "").(string),
			Username:  helpers.GetValue(cht, "username", "").(string),
		}
	}
	result.Text = helpers.GetValue(msg, "text", "").(string)
	entities, ok := msg["entities"].([]interface{})
	if ok {
		for _, entity := range entities {
			result.Entities = append(result.Entities, messageEntity{
				Type:   helpers.GetValue(entity.(map[string]interface{}), "type", "").(string),
				Offset: int(helpers.GetValue(entity.(map[string]interface{}), "offset", "").(float64)),
				Length: int(helpers.GetValue(entity.(map[string]interface{}), "length", "").(float64)),
				Url:    helpers.GetValue(entity.(map[string]interface{}), "url", "").(string),
			})
		}
	}
	return result
}

func (m message) getCommandInfo() (name string, payload string) {
	if len(m.Entities) == 0 || m.Entities[0].Type != "bot_command" {
		return "", ""
	}
	found := regexp.MustCompile(`/(?P<Command>\w+)\s(?P<Payload>\w+)`).FindStringSubmatch(m.Text)
	if found == nil || len(found) < 2 {
		return "", ""
	}
	if len(found) >= 3 {
		return found[1], found[2]
	} else {
		return found[1], ""
	}
}
