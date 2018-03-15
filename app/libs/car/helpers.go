package car

import (
	"net/url"
)

func (c Car) getApiUrl(methodName string, params map[string]string) string {
	p := url.Values{}
	for key, value := range params {
		p.Add(key, value)
	}
	return "http://api.car-online.ru/v2?get=" + methodName + "&skey=" + c.ApiKey + "&content=json&" + p.Encode()
}
