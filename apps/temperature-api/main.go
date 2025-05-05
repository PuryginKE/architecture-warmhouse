package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
    location := r.URL.Query().Get("location")
    sensorID := r.URL.Query().Get("sensorID")

    if location == "" {
        switch sensorID {
        case "1":
            location = "Living Room"
        case "2":
            location = "Bedroom"
        case "3":
            location = "Kitchen"
        default:
            location = "Unknown"
        }
    }

    if sensorID == "" {
        switch location {
        case "Living Room":
            sensorID = "1"
        case "Bedroom":
            sensorID = "2"
        case "Kitchen":
            sensorID = "3"
        default:
            sensorID = "0"
        }
    }

    rand.Seed(time.Now().UnixNano())
    randomTemp := rand.Float64() * 60 - 30
		randomTemp = math.Round(randomTemp * 100) / 100

    response := TemperatureResponse{
        Location:     location,
        SensorID:     sensorID,
        Value: randomTemp,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/temperature", temperatureHandler)
    fmt.Println("Starting server on :8081")
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}