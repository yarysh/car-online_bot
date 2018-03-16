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

	result.Id = int64(helpers.FloatDefault(msg, "message_id", 0))
	from, ok := msg["from"].(map[string]interface{})
	if ok {
		result.From = user{
			Id:           int64(helpers.FloatDefault(from, "id", 0)),
			IsBot:        helpers.BoolDefault(from, "is_bot", false),
			FirstName:    helpers.StringDefault(from, "first_name", ""),
			LastName:     helpers.StringDefault(from, "last_name", ""),
			Username:     helpers.StringDefault(from, "username", ""),
			LanguageCode: helpers.StringDefault(from, "language_code", ""),
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
