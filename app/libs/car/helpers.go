package car

func getApiUrl(methodName string, key string) string {
	return "http://api.car-online.ru/v2?get=" + methodName + "&skey=" + key + "&content=json"
}
