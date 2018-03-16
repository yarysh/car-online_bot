package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetValue(source map[string]interface{}, key string, defaultValue interface{}) interface{} {
	value, ok := source[key]
	if !ok {
		return defaultValue
	}
	return value
}

func GetJsonResponse(url string) (map[string]interface{}, error) {
	var result map[string]interface{}
	resp, err := http.Get(url)
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
	return result, nil
}
