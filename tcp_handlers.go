package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"strings"

	"github.com/plant-healthcheck-server/model"
)

// Handles incoming requests.
func (app *Application) HandleTCPRequest(conn net.Conn) error {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("error reading ", err.Error())
	}

	fmt.Println(string(buf))

	rawData := string(buf)
	sensorData, err := parseSensorData(strings.TrimRight(rawData, "\x00"))
	if err != nil {
		return fmt.Errorf("error parsing ", err.Error())
	}

	element := model.HistoryElement{
		DateStr:      time.Now().Format("2006-01-02 15:04:05"),
		Temperature:  sensorData.Data[0],
		Humidity:     sensorData.Data[1],
		SoilMoisture: sensorData.Data[2],
	}

	app.History.Push(element)

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()

	return nil
}

// aa-1;0;1.1;2.2
func parseSensorData(data string) (model.SensorData, error) {
	parsed := strings.Split(data, ";")
	if len(parsed) < 3 {
		return model.SensorData{}, errors.New("failed to parse TCP request data, size invalid")
	}

	command, err := strconv.Atoi(parsed[1])
	if err != nil {
		return model.SensorData{}, errors.New("failed to parse TCP request data, Command invalid")
	}

	sensorData := model.SensorData{
		SensorID: parsed[0],
		Command:  command,
	}

	var values []float32
	for i := 2; i < len(parsed); i++ {
		val, err := strconv.ParseFloat(parsed[i], 32)
		if err != nil {
			return model.SensorData{}, errors.New("float parsing error")
		}

		values = append(values, float32(val))
	}

	sensorData.Data = values

	return sensorData, nil
}
