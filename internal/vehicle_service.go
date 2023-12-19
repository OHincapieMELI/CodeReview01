package internal

import (
	"errors"
)

var (
	// ErrServiceVehicleNotFound is returned when no vehicle is found.
	ErrServiceVehicleNotFound = errors.New("service: vehicle not found")
	ErrServiceFieldsRequired  = errors.New("service: fields required")
)

type ErrFieldsRequired struct {
	Msg    string
	Fields map[string]string
}

func (e ErrFieldsRequired) Error() string {
	return e.Msg
}

// ServiceVehicle is the interface that wraps the basic methods for a vehicle service.
// - conections with external apis
// - business logic
type ServiceVehicle interface {
	// FindAll returns all vehicles
	FindAll() (v []Vehicle, err error)
	Save(vehicle *Vehicle) (v Vehicle, err error)
	UpdateMaxSpeed(maxSpeed int, vehicleID int) (v Vehicle, err error)
}
