package handler

import (
	"app/internal"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		length := r.URL.Query().Get("length")
		width := r.URL.Query().Get("width")
		lengths := strings.Split(length, "-")
		widths := strings.Split(width, "-")

		minLength, err := strconv.ParseFloat(lengths[0], 64)
		if err != nil {
			http.Error(w, "Erro ao converter valor.", http.StatusInternalServerError)
			return
		}
		maxLength, err := strconv.ParseFloat(lengths[1], 64)
		if err != nil {
			http.Error(w, "Erro ao converter valor.", http.StatusInternalServerError)
			return
		}
		minWidth, err := strconv.ParseFloat(widths[0], 64)
		if err != nil {
			http.Error(w, "Erro ao converter valor.", http.StatusInternalServerError)
			return
		}
		maxWidth, err := strconv.ParseFloat(widths[1], 64)
		if err != nil {
			http.Error(w, "Erro ao converter valor.", http.StatusInternalServerError)
			return
		}

		finalList, err := h.sv.GetByDimensions(minLength, maxLength, minWidth, maxWidth)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "Não foram encontrados veículos com essas dimensões.")
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range finalList {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}
