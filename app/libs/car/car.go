package car

import (
	"errors"

	"github.com/yarysh/car-online_bot/app/libs/helpers"
)

func GetStatus(key string) (Status, error) {
	status := Status{}
	if key == "" {
		return status, errors.New("empty key")
	}
	resp, err := helpers.GetJsonResponse(getApiUrl("status", key))
	if err != nil {
		return status, err
	}
	gpsData, ok := resp["gps"].(map[string]interface{})
	if ok {
		status.Gps = gps{
			Lat:   float32(helpers.GetValue(gpsData, "latitude", 0).(float64)),
			Long:  float32(helpers.GetValue(gpsData, "longitude", 0).(float64)),
			Speed: int(helpers.GetValue(gpsData, "speed", 0).(float64)),
		}
	}
	voltageData, ok := resp["voltage"].(map[string]interface{})
	if ok {
		status.Voltage = voltage{
			Main:     float32(helpers.GetValue(voltageData, "main", 0).(float64)),
			Reserved: float32(helpers.GetValue(voltageData, "reserved", 0).(float64)),
		}
	}
	return status, nil
}
