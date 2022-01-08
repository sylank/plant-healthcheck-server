package model

type SensorData struct {
	SensorID     string  `json:"sensor_id"`
	Command      int     `json:"command"`
	Temperature  float32 `json:"temperature"`
	Humidity     float32 `json:"humidity"`
	SoilMoisture float32 `json:"soil_moisture"`
}
