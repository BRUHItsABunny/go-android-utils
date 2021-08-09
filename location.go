package go_android_utils

import (
	"encoding/json"
	"strings"
)

type GPSLocation struct {
	Longitude float64
	Latitude  float64
	Altitude  float64
	Accuracy  int
	Provider  string
}

type auxGPSLocation struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Altitude  float64 `json:"altitude"`
	Accuracy  int     `json:"accuracy"`
	Provider  string  `json:"provider"`
}

func (aux *auxGPSLocation) check() {
	aux.Provider = strings.ToLower(aux.Provider)
	if aux.Provider == "random" {
		aux.Provider = randomStrSlice(AvailableLocationProviders)
		aux.Accuracy = LocationProviders[aux.Provider]
	} else {
		_, ok := LocationProviders[aux.Provider]
		if !ok {
			aux.Provider = AvailableLocationProviders[2]
			aux.Accuracy = 3
		}
	}
	if aux.Accuracy <= 0 {
		aux.Accuracy = 3
	}
}

func NewLocation(provider string, lat, lon, alt float64) *GPSLocation {
	provider = strings.ToLower(provider)
	_, ok := LocationProviders[provider]
	if !ok {
		provider = randomStrSlice(AvailableLocationProviders)
	}
	acc := LocationProviders[provider]

	return &GPSLocation{
		Longitude: lon,
		Latitude:  lat,
		Altitude:  alt,
		Accuracy:  acc,
		Provider:  provider,
	}
}

func (location *GPSLocation) MarshalJSON() ([]byte, error) {

	aux := auxGPSLocation{
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
		Altitude:  location.Altitude,
		Accuracy:  location.Accuracy,
		Provider:  location.Provider,
	}

	aux.check()

	return json.Marshal(aux)
}

func (location *GPSLocation) UnmarshalJSON(data []byte) error {
	aux := new(auxGPSLocation)
	err := json.Unmarshal(data, aux)
	if err == nil {
		aux.check()
		location.Latitude = aux.Latitude
		location.Altitude = aux.Altitude
		location.Accuracy = aux.Accuracy
		location.Longitude = aux.Longitude
		location.Provider = aux.Provider
	}
	return err
}

var (
	AvailableLocationProviders = []string{
		"gps",
		"network",
		"passive",
	}
	LocationProviders = map[string]int{
		"gps":     1,
		"network": 2,
		"passive": 3,
	}
)
