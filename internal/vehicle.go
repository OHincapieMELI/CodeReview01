package internal

// VehicleAttributes is an struct that represents the attributes of a vehicle.
type VehicleAttributes struct {
	// Brand is the brand of the vehicle.
	Brand string
	// Model is the model of the vehicle.
	Model string
	// Registration is the registration of the vehicle.
	Registration string
	// Year is the fabrication year of the vehicle.
	Year int
	// Color is the color of the vehicle.
	Color string

	// MaxSpeed is the maximum speed of the vehicle.
	MaxSpeed int
	// FuelType is the fuel type of the vehicle.
	FuelType string
	// Transmission is the transmission of the vehicle.
	Transmission string

	// Passengers is the capacity of passengers of the vehicle.
	Passengers int

	// Height is the height of the vehicle.
	Height float64
	// Width is the width of the vehicle.
	Width float64

	// Weight is the weight of the vehicle.
	Weight float64
}

// Vehicle is an struct that represents a vehicle.
type Vehicle struct {
	// ID is the unique identifier of the vehicle.
	ID int
	// Attributes is the attributes of the vehicle.
	Attributes VehicleAttributes
}

func (v Vehicle) Validate() map[string]string {
	var domainErrors = make(map[string]string)
	if v.Attributes.Brand == "" || v.Attributes.Brand == "null" {
		domainErrors["brand"] = "brand is required"
	}
	if v.Attributes.Model == "" || v.Attributes.Model == "null" {
		domainErrors["model"] = "model is required"
	}
	if v.Attributes.Registration == "" || v.Attributes.Registration == "null" {
		domainErrors["registration"] = "registration is required"
	}
	if v.Attributes.Year == 0 {
		domainErrors["year"] = "year is required"
	}
	if v.Attributes.Color == "" || v.Attributes.Color == "null" {
		domainErrors["color"] = "color is required"
	}
	if v.Attributes.MaxSpeed == 0 {
		domainErrors["max_speed"] = "max_speed is required"
	}
	if v.Attributes.FuelType == "" || v.Attributes.FuelType == "null" {
		domainErrors["fuel_type"] = "fuel_type is required"
	}
	if v.Attributes.Transmission == "" || v.Attributes.Transmission == "null" {
		domainErrors["transmission"] = "transmission is required"
	}
	if v.Attributes.Passengers == 0 {
		domainErrors["passengers"] = "passengers is required"
	}
	if v.Attributes.Height == 0 {
		domainErrors["height"] = "height is required"
	}
	if v.Attributes.Width == 0 {
		domainErrors["width"] = "width is required"
	}
	if v.Attributes.Weight == 0 {
		domainErrors["weight"] = "weight is required"
	}
	return domainErrors
}
