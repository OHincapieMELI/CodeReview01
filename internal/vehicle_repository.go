package internal

import (
	"errors"
)

var (
	// ErrRepositoryVehicleNotFound is returned when a vehicle is not found.
	ErrRepositoryVehicleNotFound = errors.New("repository: vehicle not found")
	ErrRepositoryNotDataFound    = errors.New("repository: not data found")
)

// RepositoryVehicle is the interface that wraps the basic methods for a vehicle repository.
type RepositoryVehicle interface {
	FindByID(id int) (v Vehicle, err error)
	// FindAll returns all vehicles
	FindAll() (v []Vehicle, err error)
	Save(vehicle *Vehicle) (v Vehicle, err error)

	Update(vehicle *Vehicle) (v Vehicle, err error)
}
