package car

type Status struct {
	Gps     gps
	Voltage voltage
}

type gps struct {
	Lat   float32
	Long  float32
	Speed int
}

type voltage struct {
	Main     float32
	Reserved float32
}

type Temperature struct {
	Time  int
	Value int8
}

type Telemetry struct {
	AvgSpeed   float32
	EngineTime int64
	MaxSpeed   float32
	Mileage    float64
	Stands     int64
	Ways       int64
}
