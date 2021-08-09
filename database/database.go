package database

import (
	"encoding/json"
	"errors"
	go_android_utils "github.com/BRUHItsABunny/go-android-utils"
	"math/rand"
	"strings"
	"time"
)

// Here we store a few devices and way to get them, just easy access in case you want to prototype a few devices in a library fast
// TODO: Actually add a few devices here...
var DeviceDB = map[string]string{
	// "oneplus3": "",
	"oneplus5":  "{\"android_id\":{\"id\":\"163f5f0f9a621d72\"},\"locale\":{\"language\":\"en\",\"country\":\"us\"},\"android_version\":{\"sdk\":28},\"device\":\"OnePlus5\",\"manufacturer\":\"OnePlus\",\"model\":\"ONEPLUS A5000\",\"product\":\"OnePlus5\",\"build\":\"PKQ1.180716.001\",\"type\":\"user\",\"tags\":\"release-keys\",\"rom_version\":\"2002242003\",\"dpi\":420,\"resolution_horizontal\":1080,\"resolution_vertical\":1920,\"architecture\":{\"architecture\":\"arm64\"},\"time_zone\":{\"location\":\"UTC\"}}",
	"oneplus7t": "{\"android_id\":{\"id\":\"163f5f0f9a621d72\"},\"locale\":{\"language\":\"en\",\"country\":\"us\"},\"android_version\":{\"sdk\":29},\"device\":\"OnePlus7T\",\"manufacturer\":\"OnePlus\",\"model\":\"HD1905\",\"product\":\"OnePlus7T\",\"build\":\"QKQ1.190716.003\",\"type\":\"user\",\"tags\":\"release-keys\",\"rom_version\":\"2101212100\",\"dpi\":420,\"resolution_horizontal\":1080,\"resolution_vertical\":2400,\"architecture\":{\"architecture\":\"arm64\"},\"time_zone\":{\"location\":\"UTC\"}}",
	// "oneplus9": "",
	"oneplus9pro": "{\"android_id\":{\"id\":\"163f5f0f9a621d72\"},\"locale\":{\"language\":\"en\",\"country\":\"us\"},\"android_version\":{\"sdk\":30},\"device\":\"OnePlus9Pro\",\"manufacturer\":\"OnePlus\",\"model\":\"LE2125\",\"product\":\"OnePlus9Pro\",\"build\":\"RKQ1.201105.002\",\"type\":\"user\",\"tags\":\"release-keys\",\"rom_version\":\"2105290043\",\"dpi\":600,\"resolution_horizontal\":1440,\"resolution_vertical\":3216,\"architecture\":{\"architecture\":\"arm64\"},\"time_zone\":{\"location\":\"UTC\"}}",
}

var DeviceDBKeys = []string{
	// "oneplus3",
	"oneplus5",
	"oneplus7t",
	// "oneplus9",
	"oneplus9pro",
}

func GetDBDevice(key string) (*go_android_utils.Device, bool) {
	device := new(go_android_utils.Device)
	val, found := DeviceDB[key]
	if found {
		_ = json.Unmarshal([]byte(val), device)
	}
	// Device from DB needs to be random ID
	_ = device.AndroidId.Random()
	return device, found
}

func GetRandomDevice() *go_android_utils.Device {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	device := new(go_android_utils.Device)
	val := DeviceDB[DeviceDBKeys[r.Intn(len(DeviceDBKeys))]]
	_ = json.Unmarshal([]byte(val), device)
	// Device from DB needs to be random ID
	_ = device.AndroidId.Random()
	device.GPSLocation = GetRandomDBLocation(device.Locale.GetCountryISO())
	return device
}

var LocationDB = map[string]map[string]string{
	"US": {
		"newyorkcity":  "{\"latitude\": 40.712775, \"longitude\": -74.005973, \"altitude\": 10.440, \"accuracy\": 0, \"provider\": \"random\"}",
		"losangeles":   "{\"latitude\": 34.052234, \"longitude\": -118.243685, \"altitude\": 86.854, \"accuracy\": 0, \"provider\": \"random\"}",
		"chicago":      "{\"latitude\": 41.878114, \"longitude\": -87.629798, \"altitude\": 181.513, \"accuracy\": 0, \"provider\": \"random\"}",
		"houston":      "{\"latitude\": 29.760427, \"longitude\": -95.369803, \"altitude\": 14.562, \"accuracy\": 0, \"provider\": \"random\"}",
		"washington":   "{\"latitude\": 38.907192, \"longitude\": -77.036871, \"altitude\": 22.015, \"accuracy\": 0, \"provider\": \"random\"}",
		"philadelphia": "{\"latitude\": 39.952584, \"longitude\": -75.165222, \"altitude\": 14.336, \"accuracy\": 0, \"provider\": \"random\"}",
		"miami":        "{\"latitude\": 25.761680, \"longitude\": -80.191790, \"altitude\": 0.537, \"accuracy\": 0, \"provider\": \"random\"}",
	},
	"MX": {
		"mexicocity": "{\"latitude\": 19.432608, \"longitude\": -99.133208, \"altitude\": 2229.729, \"accuracy\": 0, \"provider\": \"random\"}",
	},
	"CA": {
		"toronto": "{\"latitude\": 43.653226, \"longitude\": -79.383184, \"altitude\": 91.723, \"accuracy\": 0, \"provider\": \"random\"}",
	},
}

var AvailableCountries = []string{
	"US", "MX", "CA",
}

var AvailableCities = map[string][]string{
	"US": {
		"newyorkcity",
		"losangeles",
		"chicago",
		"houston",
		"washington",
		"philadelphia",
		"miami",
	},
	"MX": {
		"mexicocity",
	},
	"CA": {
		"toronto",
	},
}

func GetDBLocation(countryISO, city string) (*go_android_utils.GPSLocation, error) {
	countryISO = strings.ToUpper(countryISO)
	city = strings.ReplaceAll(strings.ToLower(city), " ", "")
	result := new(go_android_utils.GPSLocation)
	var err error

	_, ok := AvailableCities[countryISO]
	if ok {
		val, ok := LocationDB[countryISO][city]
		if ok {
			err = json.Unmarshal([]byte(val), result)
		} else {
			err = errors.New("city not supported")
		}
	} else {
		err = errors.New("country not supported")
	}

	return result, err
}

func GetRandomDBLocation(countryISO string) *go_android_utils.GPSLocation {
	_, ok := AvailableCities[countryISO]
	if !ok {
		countryISO = randomStrSlice(AvailableCountries)
	}
	city := randomStrSlice(AvailableCities[countryISO])
	location := LocationDB[countryISO][city]

	result := new(go_android_utils.GPSLocation)
	_ = json.Unmarshal([]byte(location), result)
	return result
}

func randomStrSlice(strSlice []string) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return strSlice[r.Intn(len(strSlice))]
}
