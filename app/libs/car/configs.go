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
