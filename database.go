package go_android_utils

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"strings"
	"time"
)

// Here we store a few devices and way to get them, just easy access in case you want to prototype a few devices in a library fast
// TODO: Actually add a few devices here...
var DeviceDB = map[string]*Device{
	// "oneplus3": "",
	"oneplus5": {
		Locale: &Locale{
			Language:   "en",
			CountryISO: "US",
		},
		Version:              AndroidVersion_V9_0,
		Device:               "OnePlus5",
		Manufacturer:         "OnePlus",
		Model:                "ONEPLUS A5000",
		Product:              "OnePlus5",
		Build:                "PKQ1.180716.001",
		Type:                 "user",
		Tags:                 "release-keys",
		IncrementalVersion:   "2002242003",
		DPI:                  420,
		ResolutionHorizontal: 1080,
		ResolutionVertical:   1920,
		SimSlots: []*SIMCard{
			{
				Imei: &IMEI{
					TAC: "86463003",
				},
			},
			{
				Imei: &IMEI{
					TAC: "86463003",
				},
			},
		},
		MacAddress: &MAC{
			OUI:     "A091A2",
			Address: "",
		},
		AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
	},
	"oneplus7t": {
		Locale: &Locale{
			Language:   "en",
			CountryISO: "US",
		},
		Version:              AndroidVersion_V10_0,
		Device:               "OnePlus7T",
		Manufacturer:         "OnePlus",
		Model:                "HD1905",
		Product:              "OnePlus7T",
		Build:                "QKQ1.190716.003",
		Type:                 "user",
		Tags:                 "release-keys",
		IncrementalVersion:   "2101212100",
		DPI:                  420,
		ResolutionHorizontal: 1080,
		ResolutionVertical:   2400,
		SimSlots: []*SIMCard{
			{
				Imei: &IMEI{
					TAC: "86789104",
				},
			},
			{
				Imei: &IMEI{
					TAC: "86789104",
				},
			},
		},
		MacAddress: &MAC{
			OUI:     "A091A2",
			Address: "",
		},
		AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
	},
	// "oneplus9": "",
	"oneplus9pro": {
		Id: nil,
		Locale: &Locale{
			Language:   "en",
			CountryISO: "US",
		},
		Version:              AndroidVersion_V11_0,
		Device:               "OnePlus9Pro",
		Manufacturer:         "OnePlus",
		Model:                "LE2125",
		Product:              "OnePlus9Pro",
		Build:                "RKQ1.201105.002",
		Type:                 "user",
		Tags:                 "release-keys",
		IncrementalVersion:   "2105290043",
		DPI:                  600,
		ResolutionHorizontal: 1440,
		ResolutionVertical:   3216,
		SimSlots: []*SIMCard{
			{
				Imei: &IMEI{
					TAC: "86381505",
				},
			},
		},
		MacAddress: &MAC{
			OUI:     "A091A2",
			Address: "",
		},
		AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
	},
}

var DeviceDBKeys = []string{
	// "oneplus3",
	"oneplus5",
	"oneplus7t",
	// "oneplus9",
	"oneplus9pro",
}

func GetDBDevice(key string) (*Device, bool) {
	device := new(Device)
	val, found := DeviceDB[key]
	if found {
		device = proto.Clone(val).(*Device)
	}
	// Device from DB needs to be random ID
	device.Id = NewAndroidID()
	device.Location = GetRandomDBLocation(device.Locale.GetCountryISO())
	if device.SimSlots == nil || len(device.SimSlots) == 0 {
		device.SimSlots = []*SIMCard{GetRandomDBSIMCard(device.Locale.GetCountryISO())}
	}
	for _, sim := range device.SimSlots {
		sim.Randomize(device.Locale.GetCountryISO())
		if sim.Imei == nil {
			sim.Imei = &IMEI{}
		}
		sim.Imei.Generate("", "")
	}
	if device.MacAddress == nil {
		device.MacAddress = new(MAC)
	}
	device.MacAddress.Generate("", false, true)
	return device, found
}

func GetRandomDevice() *Device {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	val := DeviceDB[DeviceDBKeys[r.Intn(len(DeviceDBKeys))]]
	device := proto.Clone(val).(*Device)
	// Device from DB needs to be random ID
	device.Id = NewAndroidID()
	device.Location = GetRandomDBLocation(device.Locale.GetCountryISO())
	if device.SimSlots == nil || len(device.SimSlots) == 0 {
		device.SimSlots = []*SIMCard{GetRandomDBSIMCard(device.Locale.GetCountryISO())}
	}
	for _, sim := range device.SimSlots {
		sim.Randomize(device.Locale.GetCountryISO())
		if sim.Imei == nil {
			sim.Imei = &IMEI{}
		}
		sim.Imei.Generate("", "")
	}

	if device.MacAddress == nil {
		device.MacAddress = new(MAC)
	}
	device.MacAddress.Generate("", false, true)
	return device
}

var LocationDB = map[string]map[string]*GPSLocation{
	"US": {
		"newyorkcity": &GPSLocation{
			Longitude: -74.005973,
			Latitude:  40.712775,
			Altitude:  10.440,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
		"losangeles": &GPSLocation{
			Longitude: -118.243685,
			Latitude:  34.052234,
			Altitude:  86.854,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
		"chicago": &GPSLocation{
			Longitude: -87.629798,
			Latitude:  41.878114,
			Altitude:  181.513,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
		"houston": &GPSLocation{
			Longitude: -95.369803,
			Latitude:  29.760427,
			Altitude:  14.562,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
		"washington": &GPSLocation{
			Longitude: -77.036871,
			Latitude:  38.907192,
			Altitude:  22.015,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
		"philadelphia": &GPSLocation{
			Longitude: -75.165222,
			Latitude:  39.952584,
			Altitude:  14.336,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
		"miami": &GPSLocation{
			Longitude: -80.191790,
			Latitude:  25.761680,
			Altitude:  0.537,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
	},
	"MX": {
		"mexicocity": &GPSLocation{
			Longitude: -99.133208,
			Latitude:  19.432608,
			Altitude:  2229.729,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
	},
	"CA": {
		"toronto": &GPSLocation{
			Longitude: -79.383184,
			Latitude:  43.653226,
			Altitude:  91.723,
			Provider:  LocationProvider_LocationProvider_NONE,
		},
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

func GetDBLocation(countryISO, city string) (*GPSLocation, error) {
	countryISO = strings.ToUpper(countryISO)
	city = strings.ReplaceAll(strings.ToLower(city), " ", "")
	result := new(GPSLocation)
	var err error

	_, ok := AvailableCities[countryISO]
	if ok {
		result, ok = LocationDB[countryISO][city]
		if !ok {
			err = errors.New("city not supported")
		}
	} else {
		err = errors.New("country not supported")
	}

	return result, err
}

func GetRandomDBLocation(countryISO string) *GPSLocation {
	_, ok := AvailableCities[countryISO]
	if !ok {
		countryISO = randomStrSlice(AvailableCountries)
	}
	city := randomStrSlice(AvailableCities[countryISO])
	location := LocationDB[countryISO][city]

	return proto.Clone(location).(*GPSLocation)
}
