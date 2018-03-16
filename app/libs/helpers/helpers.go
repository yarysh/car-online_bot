package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yarysh/car-online_bot/app/models"
)

func BoolDefault(source map[string]interface{}, key string, defaultValue bool) bool {
	value, ok := getValue(source, key, defaultValue).(bool)
	if ok {
		return value
	} else {
		return defaultValue
	}
}

func FloatDefault(source map[string]interface{}, key string, defaultValue float64) float64 {
	value, ok := getValue(source, key, defaultValue).(float64)
	if ok {
		return value
	} else {
		return defaultValue
	}
}

func StringDefault(source map[string]interface{}, key string, defaultValue string) string {
	value, ok := getValue(source, key, defaultValue).(string)
	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetJsonResponse(url string, logging bool) (map[string]interface{}, error) {
	var result map[string]interface{}

	rl := models.RequestLog{Url: url}
	resp, err := http.Get(url)
	if err != nil {
		if logging {
			rl.Error = err.Error()
			rl.Save()
		}
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	rl.Response = string(body)
	if err != nil {
		if logging {
			rl.Error = err.Error()
			rl.Save()
		}
		return result, err
	}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		if logging {
			rl.Error = jsonErr.Error()
			rl.Save()
		}
		return result, jsonErr
	}
	if logging {
		rl.Save()
	}
	return result, nil
}

func getValue(source map[string]interface{}, key string, defaultValue interface{}) interface{} {
	value, ok := source[key]
	if !ok {
		return defaultValue
	}
	return value
}
