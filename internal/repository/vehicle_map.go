package repository

import (
	"app/internal"
	"errors"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for index, vehicle := range r.db {
		if vehicle.Length >= minLength && vehicle.Length <= maxLength && vehicle.Width >= minWidth && vehicle.Width <= maxWidth {
			v[index] = vehicle
		}
	}
	if len(v) == 0 {
		err = errors.New("vehicle not found")
		return nil, err
	}
	return
}
