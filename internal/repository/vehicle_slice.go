package repository

import "app/internal"

// NewVehicleSlice returns a new instance of a vehicle repository in an slice.
func NewVehicleSlice(db []internal.Vehicle, lastId int) *VehicleSlice {
	return &VehicleSlice{
		db:     db,
		lastId: lastId,
	}
}

// VehicleSlice is an struct that represents a vehicle repository in an slice.
type VehicleSlice struct {
	// db is the database of vehicles.
	db []internal.Vehicle
	// lastId is the last id of the database.
	lastId int
}

// FindByID returns a vehicle by id.
func (s *VehicleSlice) FindByID(id int) (v internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	// find the vehicle by id
	for _, v := range s.db {
		if v.ID == id {
			return v, nil
		}
	}

	err = internal.ErrRepositoryNotDataFound
	return
}

// FindAll returns all vehicles
func (s *VehicleSlice) FindAll() (v []internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	// make a copy of the database
	v = make([]internal.Vehicle, len(s.db))
	copy(v, s.db)
	return
}

func (r *VehicleSlice) Save(vehicle *internal.Vehicle) (response internal.Vehicle, err error) {
	r.lastId++
	vehicle.ID = r.lastId
	r.db = append(r.db, *vehicle)
	response = *vehicle
	return
}

func (r *VehicleSlice) Update(vehicle *internal.Vehicle) (response internal.Vehicle, err error) {
	for i, v := range r.db {
		if v.ID == vehicle.ID {
			r.db[i] = *vehicle
			response = *vehicle
			return
		}
	}
	err = internal.ErrRepositoryVehicleNotFound
	return
}
