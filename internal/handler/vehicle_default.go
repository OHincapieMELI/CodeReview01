package handler

import (
	"app/internal"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// VehicleJSON is an struct that represents a vehicle in json format.
type VehicleJSON struct {
	ID int `json:"id"`
	VehiclesAttJSON
}

// VehicleJSON to save a vehicle in json format.
type VehiclesAttJSON struct {
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}

// NewVehicleDefault returns a new instance of a vehicle handler.
func NewVehicleDefault(sv internal.ServiceVehicle) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is an struct that contains handlers for vehicle.
type VehicleDefault struct {
	sv internal.ServiceVehicle
}

// GetAll returns all vehicles.
func (c *VehicleDefault) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...

		// process
		// - get all vehicles from the service
		vehicles, err := c.sv.FindAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = mapToJSON(vehicle)
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

// Save saves a vehicle.
func (c *VehicleDefault) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body VehicleJSON
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
			
		}
		vehicle := mapToDomain(body)
		vehicle, err := c.sv.Save(&vehicle)
		if err != nil {
			var fieldError *internal.ErrFieldsRequired
			if errors.As(err, &fieldError) {
				fieldsRequired := []string{}
				for _, value := range fieldError.Fields {
					fieldsRequired = append(fieldsRequired, value)
				}
				ctx.JSON(http.StatusBadRequest, map[string]any{"message": "fields required", "fields": fieldsRequired})
			} else {
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}
		ctx.JSON(http.StatusCreated, map[string]any{"message": "success to save vehicle", "data": mapToJSON(vehicle)})
	}
}

func (c *VehicleDefault) UpdateMaxSpeed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - get id from url
		id := ctx.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}
		// - get max speed from body
		var body VehicleJSON
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}
		// process
		// - update max speed
		vehicle, err := c.sv.UpdateMaxSpeed(body.MaxSpeed, idInt)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicle not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}
		// response
		// - serialize vehicle
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to update vehicle", "data": mapToJSON(vehicle)})
	}
}

func mapToJSON(vehicle internal.Vehicle) (response VehicleJSON) {
	response = VehicleJSON{
		ID: vehicle.ID,
		VehiclesAttJSON: VehiclesAttJSON{
			Brand:        vehicle.Attributes.Brand,
			Model:        vehicle.Attributes.Model,
			Registration: vehicle.Attributes.Registration,
			Year:         vehicle.Attributes.Year,
			Color:        vehicle.Attributes.Color,
			MaxSpeed:     vehicle.Attributes.MaxSpeed,
			FuelType:     vehicle.Attributes.FuelType,
			Transmission: vehicle.Attributes.Transmission,
			Passengers:   vehicle.Attributes.Passengers,
			Height:       vehicle.Attributes.Height,
			Width:        vehicle.Attributes.Width,
			Weight:       vehicle.Attributes.Weight,
		},
	}
	return
}

func mapToDomain(vehicleJSON VehicleJSON) (vehicle internal.Vehicle) {
	vehicle = internal.Vehicle{
		ID: vehicleJSON.ID,
		Attributes: internal.VehicleAttributes{
			Brand:        vehicleJSON.Brand,
			Model:        vehicleJSON.Model,
			Registration: vehicleJSON.Registration,
			Year:         vehicleJSON.Year,
			Color:        vehicleJSON.Color,
			MaxSpeed:     vehicleJSON.MaxSpeed,
			FuelType:     vehicleJSON.FuelType,
			Transmission: vehicleJSON.Transmission,
			Passengers:   vehicleJSON.Passengers,
			Height:       vehicleJSON.Height,
			Width:        vehicleJSON.Width,
			Weight:       vehicleJSON.Weight,
		},
	}
	return
}
