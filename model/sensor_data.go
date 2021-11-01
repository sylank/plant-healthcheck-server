package model

type SensorData struct {
	SensorID string
	Command  int
	Data     []float32 // 0 - Temperature, 1 - Humidity, 2 - Soil Moisture
}
