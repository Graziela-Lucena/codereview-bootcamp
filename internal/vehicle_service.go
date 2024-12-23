package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]Vehicle, err error)
}
