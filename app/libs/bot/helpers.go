package bot

import (
	"regexp"
)

func getMessage(update map[string]interface{}) message {
	result := message{}

	msg, ok := update["message"].(map[string]interface{})
	if !ok {
		return result
	}

	result.Id = int64(getValue(msg, "message_id", 0).(float64))
	from, ok := msg["from"].(map[string]interface{})
	if ok {
		result.From = user{
			Id:           int64(getValue(from, "id", 0).(float64)),
			IsBot:        getValue(from, "is_bot", false).(bool),
			FirstName:    getValue(from, "first_name", "").(string),
			LastName:     getValue(from, "last_name", "").(string),
			Username:     getValue(from, "username", "").(string),
			LanguageCode: getValue(from, "language_code", "").(string),
		}
	}
	cht, ok := msg["chat"].(map[string]interface{})
	if ok {
		result.Chat = chat{
			Id:        int64(getValue(cht, "id", 0).(float64)),
			Type:      getValue(cht, "type", "").(string),
			FirstName: getValue(cht, "first_name", "").(string),
			LastName:  getValue(cht, "last_name", "").(string),
			Username:  getValue(cht, "username", "").(string),
		}
	}
	result.Text = getValue(msg, "text", "").(string)
	entities, ok := msg["entities"].([]interface{})
	if ok {
		for _, entity := range entities {
			result.Entities = append(result.Entities, messageEntity{
				Type:   getValue(entity.(map[string]interface{}), "type", "").(string),
				Offset: int(getValue(entity.(map[string]interface{}), "offset", "").(float64)),
				Length: int(getValue(entity.(map[string]interface{}), "length", "").(float64)),
				Url:    getValue(entity.(map[string]interface{}), "url", "").(string),
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

func getValue(source map[string]interface{}, key string, defaultValue interface{}) interface{} {
	value, ok := source[key]
	if !ok {
		return defaultValue
	}
	return value
}
