package car

import (
	"strconv"

	"github.com/yarysh/car-online_bot/app/libs/helpers"
)

type Car struct {
	ApiKey string
}

func (c Car) GetStatus() (Status, error) {
	status := Status{}
	resp, err := helpers.GetJsonResponse(c.getApiUrl("status", nil))
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

func (c Car) GetTemperature(begin int, end int) ([]Temperature, error) {
	var temps []Temperature
	resp, err := helpers.GetJsonResponse(
		c.getApiUrl("temperaturelist", map[string]string{"begin": strconv.Itoa(begin), "end": strconv.Itoa(end)}),
	)
	if err != nil {
		return temps, err
	}
	tempList, ok := resp["temperatureList"].([]interface{})
	if ok {
		for _, temp := range tempList {
			temps = append(temps, Temperature{
				Time:  int(helpers.GetValue(temp.(map[string]interface{}), "eventTime", 0).(float64)),
				Value: int8(helpers.GetValue(temp.(map[string]interface{}), "value", 0).(float64)),
			})
		}
	}
	return temps, nil
}

func (c Car) GetTelemetry(begin int, end int) (Telemetry, error) {
	telemetry := Telemetry{}
	resp, err := helpers.GetJsonResponse(
		c.getApiUrl("telemetry", map[string]string{"begin": strconv.Itoa(begin), "end": strconv.Itoa(end)}),
	)
	if err != nil {
		return telemetry, err
	}
	telemetry.AvgSpeed = float32(helpers.GetValue(resp, "averageSpeed", 0).(float64))
	telemetry.EngineTime = int64(helpers.GetValue(resp, "engineTime", 0).(float64))
	telemetry.MaxSpeed = float32(helpers.GetValue(resp, "maxSpeed", 0).(float64))
	telemetry.Mileage = helpers.GetValue(resp, "mileage", 0).(float64)
	telemetry.Stands = int64(helpers.GetValue(resp, "standsCount", 0).(float64))
	telemetry.Ways = int64(helpers.GetValue(resp, "waysCount", 0).(float64))
	return telemetry, nil
}
