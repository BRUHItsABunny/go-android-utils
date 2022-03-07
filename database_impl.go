package go_android_utils

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
)

// Refactor how we store and retrieve devices, allow for easy backend switching
type AUDOptions map[string]interface{}

type AUDDatabaseIFace interface {
	GetDevice(ctx context.Context, parameters ...AUDOptions) (*Device, error)
}

var (
	ErrNoDeviceFound = errors.New("no device found")
)

// Example implementation of a JSON backed DB
type JSONAUDDatabase struct {
	Devices []*Device
}

func NewJSONAUDDevice(fileLocation string) (AUDDatabaseIFace, error) {
	f, err := os.Open(fileLocation)
	var (
		fBytes []byte
		result = &JSONAUDDatabase{Devices: []*Device{}}
	)
	if err == nil {
		fBytes, err = io.ReadAll(f)
		if err == nil {
			err = json.Unmarshal(fBytes, &result.Devices)
		}
	}

	return result, err
}

func (d *JSONAUDDatabase) GetDevice(ctx context.Context, parameters ...AUDOptions) (*Device, error) {
	// var (
	// 	params AUDOptions
	//
	// 	realMinSDK int
	// )
	// if len(parameters) > 0 {
	// 	params = parameters[0]
	// }

	// minSDK, okSDK := params["minSDK"]
	// if okSDK {
	// 	realMinSDK = minSDK.(int)
	// }

	var (
		result *Device
		err    error
	)

	for result == nil && err == nil {
		result = d.Devices[randomInt(0, len(d.Devices))]
	}

	return result, err
}
