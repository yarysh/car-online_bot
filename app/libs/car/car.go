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
			Lat:   float32(helpers.FloatDefault(gpsData, "latitude", 0)),
			Long:  float32(helpers.FloatDefault(gpsData, "longitude", 0)),
			Speed: int(helpers.FloatDefault(gpsData, "speed", 0)),
		}
	}
	voltageData, ok := resp["voltage"].(map[string]interface{})
	if ok {
		status.Voltage = voltage{
			Main:     float32(helpers.FloatDefault(voltageData, "main", 0)),
			Reserved: float32(helpers.FloatDefault(voltageData, "reserved", 0)),
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
				Time:  int(helpers.FloatDefault(temp.(map[string]interface{}), "eventTime", 0)),
				Value: int8(helpers.FloatDefault(temp.(map[string]interface{}), "value", 0)),
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
	telemetry.AvgSpeed = float32(helpers.FloatDefault(resp, "averageSpeed", 0))
	telemetry.EngineTime = int64(helpers.FloatDefault(resp, "engineTime", 0))
	telemetry.MaxSpeed = float32(helpers.FloatDefault(resp, "maxSpeed", 0))
	telemetry.Mileage = helpers.FloatDefault(resp, "mileage", 0)
	telemetry.Stands = int64(helpers.FloatDefault(resp, "standsCount", 0))
	telemetry.Ways = int64(helpers.FloatDefault(resp, "waysCount", 0))
	return telemetry, nil
}
