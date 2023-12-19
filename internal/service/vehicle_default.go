package service

import (
	"app/internal"
	"errors"
	"fmt"
)

// NewDefault returns a new instance of a vehicle service.
func NewDefault(rp internal.RepositoryVehicle) *Default {
	return &Default{rp: rp}
}

// Default is an struct that represents a vehicle service.
type Default struct {
	rp internal.RepositoryVehicle
}

// FindAll returns all vehicles.
func (s *Default) FindAll() (v []internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.FindAll()
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) Save(vehicle *internal.Vehicle) (v internal.Vehicle, err error) {
	domainErrors := vehicle.Validate()
	if len(domainErrors) > 0 {
		err = &internal.ErrFieldsRequired{Msg: internal.ErrServiceFieldsRequired.Error(), Fields: domainErrors}
		return
	}
	return s.rp.Save(vehicle)
}

func (s *Default) UpdateMaxSpeed(maxSpeed int, vehicleID int) (v internal.Vehicle, err error) {
	vehicle, err := s.rp.FindByID(vehicleID)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	vehicle.Attributes.MaxSpeed = maxSpeed
	return s.rp.Update(&vehicle)
}
