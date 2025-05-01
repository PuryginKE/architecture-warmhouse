package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"smarthome/db"
	"smarthome/models"
	"smarthome/services"
)

// // SensorHandler handles sensor-related requests
// type SensorHandler struct {
// 	DB                 *db.DB
// 	TemperatureService *services.TemperatureService
// }

// // NewSensorHandler creates a new SensorHandler
// func NewSensorHandler(db *db.DB, temperatureService *services.TemperatureService) *SensorHandler {
// 	return &SensorHandler{
// 		DB:                 db,
// 		TemperatureService: temperatureService,
// 	}
// }

// // RegisterRoutes registers the sensor routes
// func (h *SensorHandler) RegisterRoutes(router *gin.RouterGroup) {
// 	sensors := router.Group("/sensors")
// 	{
// 		sensors.GET("", h.GetSensors)
// 		sensors.GET("/:id", h.GetSensorByID)
// 		sensors.POST("", h.CreateSensor)
// 		sensors.PUT("/:id", h.UpdateSensor)
// 		sensors.DELETE("/:id", h.DeleteSensor)
// 		sensors.PATCH("/:id/value", h.UpdateSensorValue)
// 		sensors.GET("/temperature/:location", h.GetTemperatureByLocation)
// 	}
// }

// GetSensorByID handles GET /api/v1/sensors/:id
func (h *SensorHandler) GetSensorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	sensor, err := h.DB.GetSensorByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sensor not found"})
		return
	}

	// If this is a temperature sensor, fetch real-time data from the temperature API
	if sensor.Type == models.Temperature {
		tempData, err := h.TemperatureService.GetTemperatureByID(fmt.Sprintf("%d", sensor.ID))
		if err == nil {
			// Update sensor with real-time data
			sensor.Value = tempData.Value
			sensor.Status = tempData.Status
			sensor.LastUpdated = tempData.Timestamp
			log.Printf("Updated temperature data for sensor %d from external API", sensor.ID)
		} else {
			log.Printf("Failed to fetch temperature data for sensor %d: %v", sensor.ID, err)
		}
	}

	c.JSON(http.StatusOK, sensor)
}